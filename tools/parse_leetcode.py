#!/usr/bin/env python3
"""
Parses the 205 LeetCode problem .mdx files from the user's local
swiftui.uz-main checkout (data/blog/coding/Leetcode/{Easy,Medium,Hard})
into one structured JSON dataset, with every test case's arguments and
expected value pre-converted to ready-to-paste Go literal syntax.

Unlike the Abramyan set, these files only carry a Swift starter
signature (no Python), and each test case's exact positional arguments
are given directly via functionCall.Swift (e.g. "solve([2,7,11,15], 9)")
rather than needing to be inferred from a human-readable "input" string
— that's what we parse arguments from here, since it's unambiguous.
"""
import json
import re
import sys
from pathlib import Path

SRC_ROOT = Path("/Users/ai/Documents/swiftui.uz-main/data/blog/coding/Leetcode")
OUT_PATH = Path("/Users/ai/Documents/go-learn/tools/leetcode.json")

DIFFICULTY_DIRS = ["Easy", "Medium", "Hard"]

SWIFT_TO_GO = {
    "Int": "int",
    "Double": "float64",
    "Bool": "bool",
    "String": "string",
    "Character": "string",
}


def parse_frontmatter(text):
    m = re.match(r"^---\n(.*?)\n---\n", text, re.DOTALL)
    fm = {}
    if not m:
        return fm
    for line in m.group(1).split("\n"):
        line = line.strip()
        if not line or ":" not in line:
            continue
        key, _, val = line.partition(":")
        key = key.strip()
        val = val.strip()
        if (val.startswith("'") and val.endswith("'")) or (val.startswith('"') and val.endswith('"')):
            val = val[1:-1]
        fm[key] = val
    return fm


def extract_section(text, header):
    pattern = re.compile(rf"##\s*{re.escape(header)}\s*\n(.*?)(?=\n##|\n<CodeChallenge|\Z)", re.DOTALL)
    m = pattern.search(text)
    return m.group(1).strip() if m else ""


def extract_prop_template(text, prop):
    m = re.search(rf"{prop}=\{{`(.*?)`\}}", text, re.DOTALL)
    return m.group(1) if m else None


def extract_block_between(text, start_marker, end_marker):
    start = text.find(start_marker)
    if start == -1:
        return None
    start += len(start_marker)
    end = text.find(end_marker, start)
    if end == -1:
        return None
    return text[start:end]


def split_top_level_objects(s):
    """Split a JS array body of `{...}, {...}` into raw object strings,
    quote-aware so braces inside string literals don't throw off depth."""
    objs = []
    depth = 0
    cur = []
    started = False
    in_quote = None
    i = 0
    n = len(s)
    while i < n:
        ch = s[i]
        if in_quote:
            if started:
                cur.append(ch)
            if ch == "\\" and i + 1 < n:
                if started:
                    cur.append(s[i + 1])
                i += 2
                continue
            if ch == in_quote:
                in_quote = None
            i += 1
            continue
        if ch in "'\"":
            in_quote = ch
            if started:
                cur.append(ch)
            i += 1
            continue
        if ch == "{":
            depth += 1
            started = True
        if started:
            cur.append(ch)
        if ch == "}":
            depth -= 1
            if depth == 0 and started:
                objs.append("".join(cur))
                cur = []
                started = False
        i += 1
    return objs


def extract_field(obj_str, field):
    m = re.search(rf"{field}:\s*'((?:[^'\\]|\\.)*)'", obj_str)
    if m:
        return m.group(1).replace("\\'", "'")
    m = re.search(rf'{field}:\s*"((?:[^"\\]|\\.)*)"', obj_str)
    if m:
        return m.group(1).replace('\\"', '"')
    m = re.search(rf"{field}:\s*(\{{)", obj_str)
    if m:
        start = m.start(1)
        depth = 0
        for i in range(start, len(obj_str)):
            if obj_str[i] == "{":
                depth += 1
            elif obj_str[i] == "}":
                depth -= 1
                if depth == 0:
                    return obj_str[start:i + 1]
    m = re.search(rf"{field}:\s*([^,{{}}]+)", obj_str)
    if m:
        return m.group(1).strip()
    return None


def split_top_level_csv(s):
    parts = []
    depth = 0
    cur = []
    in_quote = None
    for ch in s:
        if in_quote:
            cur.append(ch)
            if ch == in_quote:
                in_quote = None
            continue
        if ch in "'\"":
            in_quote = ch
            cur.append(ch)
            continue
        if ch in "[{(":
            depth += 1
        elif ch in "]})":
            depth -= 1
        if ch == "," and depth == 0:
            parts.append("".join(cur))
            cur = []
        else:
            cur.append(ch)
    if cur:
        parts.append("".join(cur))
    return [p for p in parts if p.strip() != ""]


def swift_value_to_go(raw):
    """Convert a Swift-syntax literal (as embedded in JS test data) into
    (go_literal, inferred_go_type)."""
    raw = raw.strip()

    if raw in ("true", "false"):
        return raw, "bool"

    if raw.startswith("[") and raw.endswith("]"):
        inner = raw[1:-1].strip()
        if inner == "":
            return "[]int{}", "[]int"
        elements = split_top_level_csv(inner)
        sub_results = [swift_value_to_go(e.strip()) for e in elements]
        sub_type = sub_results[0][1] if sub_results else "int"
        go_elems = [r[0] for r in sub_results]
        go_type = f"[]{sub_type}"
        return f"{go_type}{{{', '.join(go_elems)}}}", go_type

    if (raw.startswith("'") and raw.endswith("'")) or (raw.startswith('"') and raw.endswith('"')):
        return json.dumps(raw[1:-1]), "string"

    if re.match(r"^-?\d+$", raw):
        return raw, "int"
    if re.match(r"^-?\d+\.\d+$", raw):
        return raw, "float64"

    return json.dumps(raw), "string"


def extract_call_args(call_str):
    """'solve([2, 7, 11, 15], 9)' -> '[2, 7, 11, 15], 9' (paren-balanced,
    quote-aware). Returns None if parens never balance."""
    start = call_str.find("(")
    if start == -1:
        return None
    depth = 0
    in_quote = None
    i = start
    n = len(call_str)
    while i < n:
        ch = call_str[i]
        if in_quote:
            if ch == "\\":
                i += 2
                continue
            if ch == in_quote:
                in_quote = None
            i += 1
            continue
        if ch in "'\"":
            in_quote = ch
            i += 1
            continue
        if ch == "(":
            depth += 1
        elif ch == ")":
            depth -= 1
            if depth == 0:
                return call_str[start + 1:i]
        i += 1
    return None


def parse_swift_signature(swift_code):
    """func solve(_ a: String, _ b: [Int]) -> String { -> (params, return_swift_type)
    params: list of (name, swift_type)"""
    m = re.search(r"func\s+solve\s*\((.*?)\)\s*->\s*([\[\]\w]+)\s*\{", swift_code, re.DOTALL)
    if not m:
        return [], None
    params_raw = m.group(1).strip()
    return_type = m.group(2).strip()
    params = []
    if params_raw:
        for p in split_top_level_csv(params_raw):
            p = p.strip()
            # "_ name: Type" or "name: Type" or "_ name: Type = default"
            pm = re.match(r"(?:_\s+)?(\w+)\s*:\s*([\[\]\w]+)", p)
            if pm:
                params.append((pm.group(1), pm.group(2)))
    return params, return_type


def swift_type_to_go(swift_type):
    if swift_type in SWIFT_TO_GO:
        return SWIFT_TO_GO[swift_type]
    m = re.match(r"^\[\[(\w+)\]\]$", swift_type)
    if m:
        inner = SWIFT_TO_GO.get(m.group(1), m.group(1).lower())
        return f"[][]{inner}"
    m = re.match(r"^\[(\w+)\]$", swift_type)
    if m:
        inner = SWIFT_TO_GO.get(m.group(1), m.group(1).lower())
        return f"[]{inner}"
    return None  # unknown / custom type (e.g. ListNode, TreeNode)


def main():
    problems = []
    skipped = []
    for diff_dir in DIFFICULTY_DIRS:
        d = SRC_ROOT / diff_dir
        if not d.is_dir():
            print(f"WARN: missing {d}", file=sys.stderr)
            continue
        for path in sorted(d.glob("*.mdx")):
            text = path.read_text(encoding="utf-8")
            fm = parse_frontmatter(text)
            description = extract_section(text, "📝 Vazifa tavsifi")
            example_block = extract_section(text, "Misol:")

            swift_code = extract_prop_template(text, "starterCode_swift") or ""
            params, return_swift = parse_swift_signature(swift_code)
            return_go = swift_type_to_go(return_swift) if return_swift else None
            param_go_types = [swift_type_to_go(t) for _, t in params]

            if return_go is None or any(t is None for t in param_go_types):
                skipped.append((path.stem, diff_dir, return_swift, [t for _, t in params]))
                continue

            tc_block = extract_block_between(text, "testCases={[", "]}\n  examples=")
            if tc_block is None:
                tc_block = extract_block_between(text, "testCases={[", "]}\nexamples=")
            test_cases = []
            if tc_block:
                for obj in split_top_level_objects(tc_block):
                    input_raw = extract_field(obj, "input")
                    expected_raw = extract_field(obj, "expected")
                    desc_raw = extract_field(obj, "description") or ""
                    # extract_field's own brace-depth counting (not a naive
                    # non-greedy regex) is required here: a Swift call like
                    # 'solve("()[]{}")' contains a nested, balanced {} pair
                    # inside the quoted argument itself, which a lazy
                    # `\{(.*?)\}` would truncate at.
                    fc_raw = extract_field(obj, "functionCall")
                    swift_call = None
                    if fc_raw:
                        swift_call = extract_field(fc_raw, "Swift")
                    go_args = []
                    if swift_call:
                        args_raw = extract_call_args(swift_call)
                        if args_raw is not None:
                            for tok in split_top_level_csv(args_raw):
                                lit, gt = swift_value_to_go(tok.strip())
                                go_args.append({"literal": lit, "go_type": gt})
                    expected_lit, expected_go_type = swift_value_to_go(expected_raw or "")
                    test_cases.append({
                        "input_raw": input_raw,
                        "expected_raw": expected_raw,
                        "description": desc_raw,
                        "go_args": go_args,
                        "expected_go_literal": expected_lit,
                        "expected_go_type": expected_go_type,
                    })

            problems.append({
                "id": path.stem,
                "difficulty_dir": diff_dir,
                "title": fm.get("title", path.stem),
                "difficulty": fm.get("difficulty", ""),
                "points": fm.get("points", ""),
                "estimated_time": fm.get("estimatedTime", ""),
                "description": description,
                "example_block": example_block,
                "func_name": "solve",
                "params": [{"name": n, "swift_type": t, "go_type": g} for (n, t), g in zip(params, param_go_types)],
                "return_swift_type": return_swift,
                "return_go_type": return_go,
                "test_cases": test_cases,
                "hints": [],
            })

    OUT_PATH.write_text(json.dumps(problems, indent=2, ensure_ascii=False), encoding="utf-8")
    print(f"Parsed {len(problems)} problems -> {OUT_PATH}")
    print(f"Skipped {len(skipped)} (unsupported/custom types):")
    for s in skipped:
        print(" ", s)

    bad = [p["id"] for p in problems if not p["test_cases"]]
    print(f"Problems with zero test cases: {len(bad)}")
    if bad:
        print(bad[:20])


if __name__ == "__main__":
    main()

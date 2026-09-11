#!/usr/bin/env python3
"""
Parses the 205 Abramyan problem .mdx files from the swiftui.uz blog
content into one structured JSON dataset, pre-converting every test
case's arguments/expected value into ready-to-paste Go literal syntax
so downstream lesson generation doesn't need to re-parse JS.
"""
import json
import re
import sys
from pathlib import Path

SRC_ROOT = Path("/Users/ai/Documents/swiftui.uz/data/blog/coding")
OUT_PATH = Path("/Users/ai/Documents/go-learn/tools/abramyan.json")

CATEGORY_DIRS = [
    "01.Begin", "02.Integer", "03.Boolean", "04.If", "05.Case",
    "06.For", "07.While", "08.MinMax", "09.Array", "10.Matrix", "11.String",
]


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
    """Extract a `propName={`...`}` template-literal prop's raw body."""
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
    respecting brace nesting (needed since values may contain [ ] { }).
    Quote-aware: braces inside string literals (e.g. Java's `{{1,2,3}}`
    array syntax embedded in a quoted functionCall string) don't count."""
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
            cur.append(ch) if started else None
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
    """Extract `field: <value>` from a JS object literal string, where
    <value> may be a quoted string, a bracketed array/matrix, or a bare
    number/bool. Returns the raw JS-side text of the value."""
    m = re.search(rf"{field}:\s*'((?:[^'\\]|\\.)*)'", obj_str)
    if m:
        return m.group(1).replace("\\'", "'")
    m = re.search(rf'{field}:\s*"((?:[^"\\]|\\.)*)"', obj_str)
    if m:
        return m.group(1).replace('\\"', '"')
    # bracketed array/matrix value (balance brackets)
    m = re.search(rf"{field}:\s*(\[)", obj_str)
    if m:
        start = m.start(1)
        depth = 0
        for i in range(start, len(obj_str)):
            if obj_str[i] == "[":
                depth += 1
            elif obj_str[i] == "]":
                depth -= 1
                if depth == 0:
                    return obj_str[start:i + 1]
    # bare token (number, true/false, identifier) up to comma/brace
    m = re.search(rf"{field}:\s*([^,{{}}]+)", obj_str)
    if m:
        return m.group(1).strip()
    return None


def js_value_to_go(raw):
    """Convert a raw JS literal (string content already unquoted for
    strings) into (go_literal, inferred_go_type)."""
    raw = raw.strip()
    if raw is None:
        return "nil", "interface{}"

    if raw in ("true", "false"):
        return raw, "bool"

    if raw.startswith("[") and raw.endswith("]"):
        inner = raw[1:-1].strip()
        if inner == "":
            return "[]int{}", "[]int"
        elements = split_top_level_csv(inner)
        sub_results = [js_value_to_go(e.strip()) for e in elements]
        sub_type = sub_results[0][1] if sub_results else "int"
        go_elems = [r[0] for r in sub_results]
        go_type = f"[]{sub_type}"
        return f"{go_type}{{{', '.join(go_elems)}}}", go_type

    # quoted string already unquoted by extract_field for quote-delimited
    # case, but arrays-of-strings pass raw quoted tokens through split;
    # handle both a bare numeric token and a quoted string token here.
    if (raw.startswith("'") and raw.endswith("'")) or (raw.startswith('"') and raw.endswith('"')):
        return f"{json.dumps(raw[1:-1])}", "string"

    if re.match(r"^-?\d+$", raw):
        return raw, "int"
    if re.match(r"^-?\d+\.\d+$", raw):
        return raw, "float64"

    # fallback: treat as string
    return f"{json.dumps(raw)}", "string"


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
        if ch in "[{":
            depth += 1
        elif ch in "]}":
            depth -= 1
        if ch == "," and depth == 0:
            parts.append("".join(cur))
            cur = []
        else:
            cur.append(ch)
    if cur:
        parts.append("".join(cur))
    return [p for p in parts if p.strip() != ""]


def extract_call_args(call_str):
    """'funcName(1, 2, [1,2])' -> '1, 2, [1,2]' (paren-balanced, quote-aware).
    Returns None if the call string's parens never balance (malformed)."""
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


def brackets_balanced(s):
    return s.count("[") == s.count("]") and s.count("{") == s.count("}")


def parse_input_string(input_str):
    """'d=10' -> [('d','10')]; 'M=2, N=3' -> [('M','2'),('N','3')]"""
    pairs = []
    for part in split_top_level_csv(input_str):
        if "=" not in part:
            continue
        k, _, v = part.partition("=")
        pairs.append((k.strip(), v.strip()))
    return pairs


def parse_python_signature(py_code):
    """def name(a: int, b: str) -> float:  ->  (name, [(name,pytype)], returnpytype)"""
    m = re.search(r"def\s+(\w+)\s*\((.*?)\)\s*(?:->\s*([\w\[\], ]+))?\s*:", py_code, re.DOTALL)
    if not m:
        return None, [], None
    name = m.group(1)
    params_raw = m.group(2).strip()
    ret = (m.group(3) or "").strip()
    params = []
    if params_raw:
        for p in split_top_level_csv(params_raw):
            p = p.strip()
            if ":" in p:
                pname, _, ptype = p.partition(":")
                params.append((pname.strip(), ptype.strip()))
            else:
                params.append((p, ""))
    return name, params, ret


def main():
    problems = []
    for cat in CATEGORY_DIRS:
        cat_dir = SRC_ROOT / cat
        if not cat_dir.is_dir():
            print(f"WARN: missing category dir {cat_dir}", file=sys.stderr)
            continue
        for path in sorted(cat_dir.glob("*.mdx")):
            text = path.read_text(encoding="utf-8")
            fm = parse_frontmatter(text)
            description = extract_section(text, "📝 Vazifa tavsifi")
            example_block = extract_section(text, "Misol:")

            py_starter = extract_prop_template(text, "starterCode_python")
            func_name, py_params, py_ret = parse_python_signature(py_starter or "")

            tc_block = extract_block_between(text, "testCases={[", "]}\n  examples=")
            if tc_block is None:
                tc_block = extract_block_between(text, "testCases={[", "]}\nexamples=")
            test_cases = []
            if tc_block:
                for obj in split_top_level_objects(tc_block):
                    input_raw = extract_field(obj, "input")
                    expected_raw = extract_field(obj, "expected")
                    desc_raw = extract_field(obj, "description") or ""
                    py_call = None
                    fc_m = re.search(r"functionCall:\s*\{(.*?)\}", obj, re.DOTALL)
                    if fc_m:
                        py_call = extract_field(fc_m.group(1), "python")

                    # Prefer the exact positional args from functionCall.python
                    # (authoritative — matches the declared signature's arity
                    # even when the human-readable "input" string omits
                    # parameters that default to 0, e.g. a shape-type switch
                    # problem where only the relevant params are named).
                    go_args = None
                    if py_call:
                        call_args_raw = extract_call_args(py_call)
                        if call_args_raw is not None and brackets_balanced(call_args_raw):
                            tokens = split_top_level_csv(call_args_raw)
                            if len(tokens) == len(py_params) and all(brackets_balanced(t) for t in tokens):
                                go_args = []
                                for v in tokens:
                                    lit, gt = js_value_to_go(v.strip())
                                    go_args.append({"literal": lit, "go_type": gt})
                    if go_args is None:
                        arg_pairs = parse_input_string(input_raw or "")
                        go_args = []
                        for _, v in arg_pairs:
                            lit, gt = js_value_to_go(v)
                            go_args.append({"literal": lit, "go_type": gt})
                    expected_lit, expected_go_type = js_value_to_go(expected_raw or "")
                    test_cases.append({
                        "input_raw": input_raw,
                        "expected_raw": expected_raw,
                        "description": desc_raw,
                        "python_call": py_call,
                        "go_args": go_args,
                        "expected_go_literal": expected_lit,
                        "expected_go_type": expected_go_type,
                    })

            hints_m = re.search(r"hints=\{(\[.*?\])\}", text, re.DOTALL)
            hints = []
            if hints_m:
                for h in split_top_level_csv(hints_m.group(1)[1:-1]):
                    h = h.strip()
                    if (h.startswith("'") and h.endswith("'")) or (h.startswith('"') and h.endswith('"')):
                        h = h[1:-1]
                    hints.append(h)

            problems.append({
                "id": path.stem,
                "category_dir": cat,
                "title": fm.get("title", path.stem),
                "difficulty": fm.get("difficulty", ""),
                "points": fm.get("points", ""),
                "estimated_time": fm.get("estimatedTime", ""),
                "description": description,
                "example_block": example_block,
                "func_name": func_name,
                "py_params": py_params,
                "py_return": py_ret,
                "test_cases": test_cases,
                "hints": hints,
            })

    OUT_PATH.write_text(json.dumps(problems, indent=2, ensure_ascii=False), encoding="utf-8")
    print(f"Parsed {len(problems)} problems -> {OUT_PATH}")

    # quick sanity report: any problem with zero test cases or missing func_name
    bad = [p["id"] for p in problems if not p["test_cases"] or not p["func_name"]]
    print(f"Problems with issues (no test cases or no func_name): {len(bad)}")
    if bad:
        print(bad[:20])


if __name__ == "__main__":
    main()

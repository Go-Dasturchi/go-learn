#!/usr/bin/env python3
"""
Generates go-learn lesson scaffolding (lesson.md, starter, solution
stub, tests) for all 205 Abramyan practice problems from abramyan.json,
plus the manifest.json entries to append.

Solutions are left as TODO stubs — an agent (or a human) fills in the
actual algorithm afterward; this script only handles the mechanical,
100%-derivable parts: prose, signatures, and test assertions built
directly from the site's own recorded input/expected test cases.
"""
import json
from pathlib import Path

DATA_PATH = Path("/Users/ai/Documents/go-learn/tools/abramyan.json")
LESSONS_ROOT = Path("/Users/ai/Documents/go-learn/lessons")
MANIFEST_FRAGMENT_PATH = Path("/Users/ai/Documents/go-learn/tools/abramyan_manifest_fragment.json")

ZERO_VALUES = {
    "int": "0",
    "float64": "0",
    "string": '""',
    "bool": "false",
    "[]int": "[]int{}",
    "[]string": "[]string{}",
    "[][]int": "[][]int{}",
    "[][]string": "[][]string{}",
}


def subgroup_name(category_dir):
    # "01.Begin" -> "Begin"
    return category_dir.split(".", 1)[1] if "." in category_dir else category_dir


def widen_types(types):
    """Pick a single Go type that fits every test case's value at this
    position. int and float64 both appearing (e.g. solve(5) and
    solve(3.5) for the same parameter) means the true type is float64 —
    an untyped int literal converts to float64 automatically at the
    call site, so widening never breaks the existing example calls."""
    uniq = set(types)
    if len(uniq) == 1:
        return next(iter(uniq))
    if uniq <= {"int", "float64"}:
        return "float64"
    if "string" in uniq:
        # A handful of source problems format their result as text in
        # some test cases (e.g. converting to base 16 can yield "F") but
        # a plain digit string in others ("17") gets misdetected as int
        # since the JS source quotes both alike — string is the only
        # type that can hold every observed value, so prefer it whenever
        # it appears in the mix at all.
        return "string"
    return types[-1]


def go_signature(problem):
    param_names = [p[0] for p in problem["py_params"]]
    test_cases = problem["test_cases"]
    param_types = []
    for i in range(len(param_names)):
        types_at_i = [tc["go_args"][i]["go_type"] for tc in test_cases]
        param_types.append(widen_types(types_at_i))
    if len(param_names) != len(param_types):
        raise ValueError(f"{problem['id']}: param name/type count mismatch")
    params = ", ".join(f"{n} {t}" for n, t in zip(param_names, param_types))
    return_type = widen_types([tc["expected_go_type"] for tc in test_cases])
    return param_names, param_types, return_type, params


def needs_reflect(go_type):
    return go_type.startswith("[]") or go_type.startswith("[][]")


def gen_lesson_md(problem, func_name, params, return_type):
    title = problem["title"]
    difficulty = problem["difficulty"]
    points = problem["points"]
    est = problem["estimated_time"]
    description = problem["description"] or problem.get("title", "")
    example_block = problem["example_block"]
    hints = problem["hints"] or ["Masala shartini qayta o'qib, kerakli formulani aniqlang."]

    hints_md = "\n".join(f"{i+1}. {h}" for i, h in enumerate(hints))

    example_section = ""
    if example_block:
        example_section = f"## EXAMPLE\n\n{example_block}\n\n"

    return f"""# {title}

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, {subgroup_name(problem['category_dir'])} bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **{difficulty}** · Ball: **{points}** · Taxminiy vaqt: **{est}**

{example_section}## TASK

{description}

Quyidagi funksiyani to'ldiring:

```go
func {func_name}({params}) {return_type} {{
    // ...
}}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

{hints_md}
"""


def gen_starter(func_name, params, return_type, example_call_args):
    zero = ZERO_VALUES.get(return_type, "0")
    return f"""package main

import "fmt"

func {func_name}({params}) {return_type} {{
	// TODO: masalani yeching
	return {zero}
}}

func main() {{
	fmt.Println({func_name}({example_call_args}))
}}
"""


def gen_solution_stub(func_name, params, return_type, example_call_args):
    zero = ZERO_VALUES.get(return_type, "0")
    call_args = example_call_args
    return f"""package main

import "fmt"

func {func_name}({params}) {return_type} {{
	// TODO(solver): implement — see lesson.md TASK for the problem statement
	return {zero}
}}

func main() {{
	fmt.Println({func_name}({call_args}))
}}
"""


def gen_test(func_name, return_type, test_cases):
    uses_reflect = needs_reflect(return_type)
    uses_math = return_type == "float64"

    imports = ['"testing"']
    if uses_reflect:
        imports.insert(0, '"reflect"')
    if uses_math:
        imports.insert(0, '"math"')

    tests = []
    for i, tc in enumerate(test_cases, start=1):
        args = ", ".join(a["literal"] for a in tc["go_args"])
        want = tc["expected_go_literal"]
        fname = f"Test{func_name[0].upper()}{func_name[1:]}Case{i}"
        desc = tc.get("input_raw", "")
        # Declare `want` with an explicit type annotation rather than
        # `:=` — a literal like `3` would otherwise infer as `int` even
        # when the function returns float64 (e.g. a median that happens
        # to land on a whole number in some test cases), causing a
        # got-want type mismatch at compile time.
        if uses_reflect:
            body = f"""func {fname}(t *testing.T) {{
	got := {func_name}({args})
	var want {return_type} = {want}
	if !reflect.DeepEqual(got, want) {{
		t.Errorf("{func_name}({desc}):\\nExpected:\\n%v\\n\\nGot:\\n%v", want, got)
	}}
}}"""
        elif uses_math:
            body = f"""func {fname}(t *testing.T) {{
	got := {func_name}({args})
	var want {return_type} = {want}
	if math.Abs(got-want) > 1e-6 {{
		t.Errorf("{func_name}({desc}):\\nExpected:\\n%v\\n\\nGot:\\n%v", want, got)
	}}
}}"""
        else:
            body = f"""func {fname}(t *testing.T) {{
	got := {func_name}({args})
	var want {return_type} = {want}
	if got != want {{
		t.Errorf("{func_name}({desc}):\\nExpected:\\n%v\\n\\nGot:\\n%v", want, got)
	}}
}}"""
        tests.append(body)

    imports_block = "\n\t".join(imports)
    return f"""package main

import (
	{imports_block}
)

{chr(10).join('' if not t else t for t in tests)}
""".replace("\n\n\n", "\n\n")


def main():
    problems = json.loads(DATA_PATH.read_text(encoding="utf-8"))
    manifest_entries = []
    errors = []

    for problem in problems:
        pid = problem["id"]
        try:
            _, _, return_type, params = go_signature(problem)
            func_name = problem["func_name"]
            example_call_args = ", ".join(a["literal"] for a in problem["test_cases"][0]["go_args"])

            lesson_dir = LESSONS_ROOT / pid
            (lesson_dir / "starter").mkdir(parents=True, exist_ok=True)
            (lesson_dir / "solution").mkdir(parents=True, exist_ok=True)
            (lesson_dir / "tests").mkdir(parents=True, exist_ok=True)

            (lesson_dir / "lesson.md").write_text(
                gen_lesson_md(problem, func_name, params, return_type), encoding="utf-8"
            )
            (lesson_dir / "starter" / "main.go.tmpl").write_text(
                gen_starter(func_name, params, return_type, example_call_args), encoding="utf-8"
            )
            (lesson_dir / "solution" / "main.go.tmpl").write_text(
                gen_solution_stub(func_name, params, return_type, example_call_args), encoding="utf-8"
            )
            (lesson_dir / "tests" / "main_test.go.tmpl").write_text(
                gen_test(func_name, return_type, problem["test_cases"]), encoding="utf-8"
            )

            manifest_entries.append({
                "id": pid,
                "title": problem["title"],
                "category": "Problems",
                "group": "Abramyan",
                "subgroup": subgroup_name(problem["category_dir"]),
            })
        except Exception as e:
            errors.append((pid, str(e)))

    MANIFEST_FRAGMENT_PATH.write_text(json.dumps(manifest_entries, indent=2, ensure_ascii=False), encoding="utf-8")
    print(f"Generated {len(manifest_entries)} lesson scaffolds")
    print(f"Manifest fragment -> {MANIFEST_FRAGMENT_PATH}")
    if errors:
        print(f"ERRORS ({len(errors)}):")
        for pid, e in errors:
            print(f"  {pid}: {e}")


if __name__ == "__main__":
    main()

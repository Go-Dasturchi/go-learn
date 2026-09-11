#!/usr/bin/env python3
"""
Generates go-learn lesson scaffolding (lesson.md, starter, solution
stub, tests) for all 205 LeetCode problems from leetcode.json, plus
the manifest.json entries to append. Solutions are left as TODO
stubs — an agent (or a human) fills in the actual algorithm afterward.
"""
import json
from pathlib import Path

DATA_PATH = Path("/Users/ai/Documents/go-learn/tools/leetcode.json")
LESSONS_ROOT = Path("/Users/ai/Documents/go-learn/lessons")
MANIFEST_FRAGMENT_PATH = Path("/Users/ai/Documents/go-learn/tools/leetcode_manifest_fragment.json")

ZERO_VALUES = {
    "int": "0",
    "float64": "0",
    "string": '""',
    "bool": "false",
    "[]int": "[]int{}",
    "[]string": "[]string{}",
    "[]float64": "[]float64{}",
    "[]bool": "[]bool{}",
    "[][]int": "[][]int{}",
    "[][]string": "[][]string{}",
}


def needs_reflect(go_type):
    return go_type.startswith("[]")


def fix_literal(literal, target_type):
    """swift_value_to_go can't know a field's declared type from an empty
    Swift array literal alone, so it always guesses "[]int{}"/"[]int" for
    `[]` — wrong whenever the field is actually []string, [][]int, etc.
    Correct it here now that we know the real target type."""
    if literal == "[]int{}" and target_type != "[]int" and target_type in ZERO_VALUES:
        return ZERO_VALUES[target_type]
    return literal


def go_signature(problem):
    params = ", ".join(f"{p['name']} {p['go_type']}" for p in problem["params"])
    return params, problem["return_go_type"]


def gen_lesson_md(problem, func_name, params, return_type):
    title = problem["title"]
    difficulty = problem["difficulty"]
    points = problem["points"]
    est = problem["estimated_time"]
    description = problem["description"] or title
    example_block = problem["example_block"]

    example_section = ""
    if example_block:
        example_section = f"## EXAMPLE\n\n{example_block}\n\n"

    return f"""# {title}

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, {problem['difficulty_dir']} daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **{difficulty}** ({problem['difficulty_dir']}) · Ball: **{points}** · Taxminiy vaqt: **{est}**

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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
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
    return f"""package main

import "fmt"

func {func_name}({params}) {return_type} {{
	// TODO(solver): implement — see lesson.md TASK for the problem statement
	return {zero}
}}

func main() {{
	fmt.Println({func_name}({example_call_args}))
}}
"""


def gen_test(func_name, return_type, test_cases, params=None):
    uses_reflect = needs_reflect(return_type)
    uses_math = return_type == "float64"

    imports = ['"testing"']
    if uses_reflect:
        imports.insert(0, '"reflect"')
    if uses_math:
        imports.insert(0, '"math"')

    tests = []
    for i, tc in enumerate(test_cases, start=1):
        if params:
            args = ", ".join(
                fix_literal(a["literal"], params[j]["go_type"])
                for j, a in enumerate(tc["go_args"])
            )
        else:
            args = ", ".join(a["literal"] for a in tc["go_args"])
        want = fix_literal(tc["expected_go_literal"], return_type)
        fname = f"Test{func_name[0].upper()}{func_name[1:]}Case{i}"
        raw_desc = tc.get("input_raw", "") or ""
        desc = raw_desc.replace("\\", "\\\\").replace('"', '\\"').replace("\n", "\\n")
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

{chr(10).join(tests)}
"""


def main():
    problems = json.loads(DATA_PATH.read_text(encoding="utf-8"))
    manifest_entries = []
    errors = []

    for problem in problems:
        pid = problem["id"]
        try:
            params, return_type = go_signature(problem)
            raw_params = problem["params"]
            func_name = problem["func_name"]
            example_call_args = ", ".join(
                fix_literal(a["literal"], raw_params[j]["go_type"])
                for j, a in enumerate(problem["test_cases"][0]["go_args"])
            )

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
                gen_test(func_name, return_type, problem["test_cases"], raw_params), encoding="utf-8"
            )

            manifest_entries.append({
                "id": pid,
                "title": problem["title"],
                "category": "Problems",
                "group": "Leetcode",
                "subgroup": problem["difficulty_dir"],
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

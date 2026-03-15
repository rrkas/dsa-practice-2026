from pathlib import Path
import sys, os
import os.path


def check_dir(platform: str, prob_name: str):
    project_dir = Path(f"platforms/{platform}/{prob_name}")

    problem_pics = sorted((project_dir / "problem").glob("*.png"))
    test_cases_pics = sorted((project_dir / "test-cases").glob("*.png"))
    solution_files = sorted((project_dir / "solutions").glob("*.*"))

    valid_soln_exts = set()

    for e in solution_files:
        if os.stat(e).st_size == 0:
            return False, sorted(valid_soln_exts), f"Solution file `{e.name}` is empty"
        else:
            valid_soln_exts.add(os.path.splitext(e)[1][1:])

    if len(problem_pics) == 0:
        return False, sorted(valid_soln_exts), "Missing Problem Pics"

    if len(test_cases_pics) == 0:
        return False, sorted(valid_soln_exts), "Missing test cases pics"

    if len(solution_files) == 0:
        return False, sorted(valid_soln_exts), "Missing Solution Files"

    return True, sorted(valid_soln_exts), None


if __name__ == "__main__":
    platform = sys.argv[1]
    prob_name = sys.argv[2]

    print(check_dir(platform, prob_name))

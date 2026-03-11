from pathlib import Path
import sys, os


def check_dir(platform: str, prob_name: str):
    project_dir = Path(f"{platform}/{prob_name}")
    problem_pics = sorted((project_dir / "problem").glob("*.png"))
    test_cases_pics = sorted((project_dir / "test-cases").glob("*.png"))
    solution_files = sorted((project_dir / "solutions").glob("*.*"))

    if len(problem_pics) == 0:
        return False, "Missing Problem Pics"

    if len(test_cases_pics) == 0:
        return False, "Missing test cases pics"

    if len(solution_files) == 0:
        return False, "Missing Solution Files"

    for e in solution_files:
        if os.stat(e).st_size == 0:
            return False, f"Solution file `{e.name}` is empty"

    return True, None


if __name__ == "__main__":
    platform = sys.argv[1]
    prob_name = sys.argv[2]

    print(check_dir(platform, prob_name))

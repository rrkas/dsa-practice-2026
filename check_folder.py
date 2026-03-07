from pathlib import Path
import sys, os


def check_dir(platform: str, prob_name: str):
    project_dir = Path(f"{platform}/{prob_name}")
    problem_pics = sorted((project_dir / "problem").glob("*.png"))
    test_cases_pics = sorted((project_dir / "test-cases").glob("*.png"))
    solution_files = sorted((project_dir / "solutions").glob("*.*"))

    return (
        len(problem_pics) > 0
        and len(test_cases_pics) > 0
        and len(solution_files) > 0
        and all(os.stat(e).st_size > 0 for e in solution_files)
    )


if __name__ == "__main__":
    platform = sys.argv[1]
    prob_name = sys.argv[2]

    print(check_dir(platform, prob_name))

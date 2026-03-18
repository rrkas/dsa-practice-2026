from pathlib import Path
from _platform_score_level import get_platform_score_level
import sys, os
import os.path

import pandas as pd


class Record:
    def __init__(self, platform: str, prob_name: str):
        self.platform = platform
        self.prob_name = prob_name

        self.project_dir = Path(f"platforms/{platform}/{prob_name}")
        self.problem_pics = sorted((self.project_dir / "problem").glob("*.png"))
        self.test_cases = sorted((self.project_dir / "test-cases").glob("*.*"))
        self.solution_files = sorted((self.project_dir / "solutions").glob("*.*"))

        self.issue = None
        self.valid_soln_exts = set()

        self.difficulty_score = None
        self.difficulty_level = None

        self.validate()

    def validate(self):
        self.difficulty_score, self.difficulty_level = get_platform_score_level(
            self.platform, self.prob_name
        )

        for e in self.solution_files:
            if os.stat(e).st_size == 0:
                self.issue = f"Solution file `{e.name}` is empty"
                return False
            else:
                self.valid_soln_exts.add(os.path.splitext(e)[1][1:])

        if len(self.problem_pics) == 0:
            self.issue = "Missing Problem Pics"
            return False

        if len(self.test_cases) == 0:
            self.issue = "Missing test cases pics"
            return False

        if len(self.solution_files) == 0:
            self.issue = "Missing Solution Files"
            return False

        return True

    def to_dict(self):
        return {
            "PLATFORM": self.platform,
            "PROB NAME": self.prob_name,
            "DIFF_SCORE": self.difficulty_score,
            "DIFF_LEVEL": self.difficulty_level,
            "PROB_STMNT_PICS": len(self.problem_pics),
            "TEST_CASES": len(self.test_cases),
            "SOLN_EXTS": sorted(self.valid_soln_exts),
            "STATUS": "PASS" if self.issue is None else "FAIL",
            "ISSUE": self.issue or "",
        }

    def to_df(self):
        return pd.DataFrame(self.to_dict())


if __name__ == "__main__":
    platform = sys.argv[1]
    prob_name = sys.argv[2]

    print(Record(platform, prob_name).to_df())

import os
from pathlib import Path

import pandas as pd


class Record:
    def __init__(self, platform: str, prob_name: str):
        self.platform: str = platform
        self.prob_name: str = prob_name

        self.project_dir = Path(f"platforms/{platform}/{prob_name}")
        self.problem_pics = sorted((self.project_dir / "problem").glob("*.png"))
        self.test_cases = sorted((self.project_dir / "test-cases").glob("*.*"))
        self.solution_files = sorted((self.project_dir / "solutions").glob("*.*"))

        self.issue: str = None
        self.valid_soln_exts = set()

        self.difficulty_score: float = None
        self.difficulty_level: str = None
        self.prob_link: str = self._get_prob_link(self.platform, self.prob_name)

        self._validate()

    def _validate(self):
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

    def _get_prob_link(self, platform: str, prob_name: str):
        match platform:
            case "codechef":
                return ""

            case "codewars":
                prob_name = prob_name.split("--")[1]
                return f"https://www.codewars.com/kata/{prob_name}"

            case "hacker-earth":
                return ""

            case "hacker-rank":
                prob_name = prob_name.split("--")[-1]
                return f"https://www.hackerrank.com/challenges/{prob_name}/problem?isFullScreen=false"

            case "interviewbit":
                prob_name = prob_name.split("--")[-1]
                return f"https://www.interviewbit.com/problems/{prob_name}/"

            case "leetcode":
                prob_name = prob_name.split("--")[-1]
                return f"https://leetcode.com/problems/{prob_name}/description/"

            case _:
                raise ValueError(
                    f"Problem Link not implemented for platform: `{platform}`"
                )

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
            "PROB_LINK": self.prob_link,
            "ISSUE": self.issue or "",
        }

    def to_df(self):
        return pd.DataFrame(self.to_dict())


PROB_NAME_PARTS = {
    "codechef": 3,
    "codewars": 3,
    "hacker-earth": 2,
    "hacker-rank": 3,
    "interviewbit": 2,
    "leetcode": 2,
}


def get_platform_score_level(platform: str, prob_name: str):
    prob_name_parts = prob_name.split("--")
    score, level = None, None

    # SCORES
    if platform in ["codechef", "codewars"]:
        if len(prob_name_parts) < PROB_NAME_PARTS[platform]:
            raise NameError(
                f"{platform}/{prob_name} has < {PROB_NAME_PARTS[platform]} score or difficulty level in name (--)!"
            )

        score = int(prob_name_parts[0])

    # LEVELS
    elif platform in ["hacker-earth", "leetcode", "interviewbit", "machinehack"]:
        if len(prob_name_parts) < PROB_NAME_PARTS[platform]:
            raise NameError(
                f"{platform}/{prob_name} has < {PROB_NAME_PARTS[platform]} score or difficulty level in name (--)!"
            )

        level = prob_name_parts[0]

    # SCORE & LEVEL
    elif platform in ["hacker-rank"]:
        if len(prob_name_parts) < PROB_NAME_PARTS[platform]:
            raise NameError(
                f"{platform}/{prob_name} has < {PROB_NAME_PARTS[platform]} score or difficulty level in name (--)!"
            )

        level = prob_name_parts[0]
        score = int(prob_name_parts[1])

    else:
        raise NotImplementedError(
            f"score-level logic not implemented for platform `{platform}`!"
        )

    if level:
        match level:
            # normal cases
            case "E":
                level = "EASY"
            case "M":
                level = "MEDIUM"
            case "H":
                level = "HARD"

            # hacker rank
            case "A":
                level = "ADVANCED"

            case _:
                raise NotImplementedError(f"Difficulty Level `{level}` is not handled!")

    assert (
        score is not None or level is not None
    ), f"Either score or difficulty level is needed, found neither!!"

    return score, level

from pathlib import Path
from check_folder import check_dir
from loguru import logger
import pandas as pd

platforms = [
    "codechef",
    "exercism",
    "hacker-earth",
    "hacker-rank",
    "interviewbit",
    "leetcode",
]

recs = []

for platform in sorted(platforms):
    for prob_name in sorted(Path(platform).glob("*/")):
        res = check_dir(platform, prob_name.name)
        # log_fn = logger.info if res else logger.error
        # log_fn(f"{platform} {prob_name.name} {res}")
        recs.append(
            {
                "platform": platform,
                "prob name": prob_name.name,
                "status": "PASS" if res else "FAIL",
            }
        )

print(pd.DataFrame(recs).to_string())

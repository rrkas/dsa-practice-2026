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
        res, issue = check_dir(platform, prob_name.name)
        # log_fn = logger.info if res else logger.error
        # log_fn(f"{platform} {prob_name.name} {res}")
        recs.append(
            {
                "platform": platform,
                "prob name": prob_name.name,
                "status": "PASS" if res else "FAIL",
                "issue": issue or "",
            }
        )


df = pd.DataFrame(recs)
print(df[df["status"] != "PASS"].to_string())

print()
print("=-" * 50)
print()

summary = []
for (pl,), tdf in df.groupby(["platform"]):
    tot = len(tdf)
    passed = len(tdf[tdf["status"] == "PASS"])
    summary.append({"platform": pl, "passed": f"{passed}/{tot}"})

# print(summary)
print(pd.DataFrame(summary).to_string())

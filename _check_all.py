from pathlib import Path
from _check_folder import check_dir
from loguru import logger
import pandas as pd

recs = []

for platform in sorted(Path("platforms").glob("*/")):
    for prob_name in sorted(Path(platform).glob("*/")):
        res, issue = check_dir(platform, prob_name.name)
        # log_fn = logger.info if res else logger.error
        # log_fn(f"{platform} {prob_name.name} {res}")
        recs.append(
            {
                "PLATFORM": platform.name,
                "PROB NAME": prob_name.name,
                "STATUS": "PASS" if res else "FAIL",
                "ISSUE": issue or "",
            }
        )


df = pd.DataFrame(recs)
print(df[df["STATUS"] != "PASS"].to_string())

print()
print("-" * 80)
print()

summary = []
for (pl,), tdf in df.groupby(["PLATFORM"]):
    tot = len(tdf)
    passed = len(tdf[tdf["STATUS"] == "PASS"])
    summary.append({"PLATFORM": pl, "PASSED": f"{passed}/{tot}"})

# print(summary)
print(pd.DataFrame(summary).to_string())

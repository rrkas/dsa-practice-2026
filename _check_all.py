from pathlib import Path
from _check_folder import check_dir
import pandas as pd

recs = []

for platform in sorted(Path("platforms").glob("*/")):
    for prob_name in sorted(platform.glob("*/")):
        res, soln_exts, issue = check_dir(platform.name, prob_name.name)
        recs.append(
            {
                "PLATFORM": platform.name,
                "PROB NAME": prob_name.name,
                "SOLN_EXTS": soln_exts,
                "STATUS": "PASS" if res else "FAIL",
                "ISSUE": issue or "",
            }
        )


df = pd.DataFrame(recs)
df.to_csv("status.csv")
print(df[df["STATUS"] != "PASS"].to_string(justify="left"))

print()
print("-" * 80)
print()

summary = []
for (pl,), tdf in df.groupby(["PLATFORM"]):
    tot = len(tdf)
    passed = len(tdf[tdf["STATUS"] == "PASS"])
    summary.append(
        {
            "PASSED": f"{passed}/{tot}",
            "PLATFORM": pl,
        }
    )

print(pd.DataFrame(summary).to_string())

from pathlib import Path
from _check_folder import Record
import pandas as pd, json
from collections import Counter

recs = []

for platform in sorted(Path("platforms").glob("*/")):
    for prob_name in sorted(platform.glob("*/")):
        rec = Record(platform.name, prob_name.name)
        recs.append(rec.to_dict())


df = pd.DataFrame(recs)
df.to_csv("status.csv")
issue_df = df[df["STATUS"] != "PASS"]

if len(issue_df) > 0:
    print(issue_df.to_string())

    print()
    print("-" * 80)
    print()

summary = []
tot_probs = 0
tot_pass = 0
for (pl,), tdf in df.groupby(["PLATFORM"]):
    tot = len(tdf)
    tot_probs += tot
    passed = len(tdf[tdf["STATUS"] == "PASS"])
    tot_pass += passed
    summary.append(
        {
            "PLATFORM": pl,
            "PASSED": f"{passed}/{tot}",
            "DIFF_LEVELS": (
                dict(**Counter(tdf["DIFF_LEVEL"]))
                if tdf["DIFF_SCORE"].sum() == 0
                else ""
            ),
            "TOT_DIFF_SCORE": (
                tdf["DIFF_SCORE"].sum() if tdf["DIFF_SCORE"].sum() != 0 else ""
            ),
        }
    )

summary.append(
    {
        "PLATFORM": "TOTAL ==>",
        "PASSED": f"{tot_pass}/{tot_probs}",
        "DIFF_LEVELS": "",
        "TOT_DIFF_SCORE": "",
    }
)

print(pd.DataFrame(summary).to_string(index=0))

from pathlib import Path
from _check_folder import Record
import pandas as pd

recs = []

for platform in sorted(Path("platforms").glob("*/")):
    for prob_name in sorted(platform.glob("*/")):
        rec = Record(platform.name, prob_name.name)
        recs.append(rec.to_dict())


df = pd.DataFrame(recs)
df.to_csv("status.csv")
print(df[df["STATUS"] != "PASS"].to_string())

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
        }
    )

summary.append(
    {
        "PLATFORM": "TOTAL ==>",
        "PASSED": f"{tot_pass}/{tot_probs}",
    }
)

print(pd.DataFrame(summary).to_string(index=0))

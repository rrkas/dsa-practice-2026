from pathlib import Path
from check_folder import check_dir

platforms = [
    "codechef",
    "exercism",
    "hacker-earth",
    "hacker-rank",
    "interviewbit",
    "leetcode",
]

for platform in platforms:
    for prob_name in sorted(Path(platform).glob("*/")):
        print(platform, prob_name.name, check_dir(platform, prob_name.name))

from pathlib import Path
from check_folder import check_dir
from loguru import logger

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
        res = check_dir(platform, prob_name.name)
        log_fn = logger.info if res else logger.error
        log_fn(f"{platform} {prob_name.name} {res}")

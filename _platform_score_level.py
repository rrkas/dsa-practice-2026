def get_platform_score_level(platform: str, prob_name: str):
    prob_name_parts = prob_name.split("--")
    if len(prob_name_parts) == 1:
        raise NameError(
            f"{platform}/{prob_name} has no score or difficulty level in name (--)!"
        )

    score, level = None, None

    # SCORES
    if platform in ["codechef", "codewars", "hacker-rank"]:
        score = int(prob_name_parts[0])

    # LEVELS
    elif platform in ["hacker-earth", "leetcode", "interviewbit", "machinehack"]:
        level = prob_name_parts[0]

    # SCORE & LEVEL
    elif platform in []:
        pass

    else:
        raise NotImplementedError(
            f"score-level logic not implemented for platform `{platform}`!"
        )

    if level:
        match level:
            case "E":
                level = "EASY"
            case "M":
                level = "MEDIUM"
            case "H":
                level = "HARD"
            case _:
                raise NotImplementedError(f"Difficulty Level `{level}` is not handled!")

    assert (
        score is not None or level is not None
    ), f"Either score or difficulty level is needed, found neither!!"

    return score, level

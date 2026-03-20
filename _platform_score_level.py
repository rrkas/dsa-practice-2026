def get_platform_score_level(platform: str, prob_name: str):
    prob_name_parts = prob_name.split("--")
    score, level = None, None

    # SCORES
    if platform in ["codechef", "codewars"]:
        if len(prob_name_parts) < 2:
            raise NameError(
                f"{platform}/{prob_name} has < 2 score or difficulty level in name (--)!"
            )

        score = int(prob_name_parts[0])

    # LEVELS
    elif platform in ["hacker-earth", "leetcode", "interviewbit", "machinehack"]:
        if len(prob_name_parts) < 2:
            raise NameError(
                f"{platform}/{prob_name} has < 2 score or difficulty level in name (--)!"
            )

        level = prob_name_parts[0]

    # SCORE & LEVEL
    elif platform in ["hacker-rank"]:
        if len(prob_name_parts) < 3:
            raise NameError(
                f"{platform}/{prob_name} has < 3 score or difficulty level in name (--)!"
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

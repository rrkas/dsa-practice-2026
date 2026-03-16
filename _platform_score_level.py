def get_platform_score_level(platform: str, prob_name: str):
    prob_name_parts = prob_name.split("--")
    if len(prob_name_parts) == 1:
        raise NameError(
            f"{platform}/{prob_name} has no score or difficulty level in name (--)"
        )

    score, level = None, None
    match platform:
        case "codechef":
            score = int(prob_name_parts[0])
        case "codewars":
            score = int(prob_name_parts[0])
        case "hacker-rank":
            score = int(prob_name_parts[0])
        case "hacker-earth":
            level = prob_name_parts[0]
        case "leetcode":
            level = prob_name_parts[0]
        case _:
            raise NotImplementedError(
                f"score-level logic not implemented for platform `{platform}`"
            )

    assert (
        score is not None or level is not None
    ), f"Either score or difficulty level is needed, found neither!!"

    return score, level

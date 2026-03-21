from _platform_score_level import Record
import sys

if __name__ == "__main__":
    platform = sys.argv[1]
    prob_name = sys.argv[2]

    print(Record(platform, prob_name).to_df())

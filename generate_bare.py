from pathlib import Path
import sys, os


def generate_new(platform: str, prob_name: str):
    platform_dir = (Path("platforms") / platform).resolve()
    os.makedirs(platform_dir, exist_ok=True)

    prob_dir = platform_dir / prob_name
    os.makedirs(prob_dir)

    sub_dirs = ["problem", "test-cases", "solutions"]
    for s in sub_dirs:
        os.makedirs(prob_dir / s)

    sol_files = ["solution.go", "solution.py"]
    for s in sol_files:
        fp = prob_dir / "solutions" / s
        with open(fp, "w"):
            pass

        os.system(f"code {fp}")

    os.system(f"nautilus {prob_dir}")


if __name__ == "__main__":
    generate_new(sys.argv[1], sys.argv[2])

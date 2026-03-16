black --fast .
uv run _check_all.py 2>&1 | tee status.txt

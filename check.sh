(
    printf "==========================================================================================\n\n"
    echo $(date)
    printf "\n"
    uv run _check_all.py  
    printf "\n"
)  2>&1 | tee -a status.txt

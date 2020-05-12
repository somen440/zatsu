CSV_1=$0
CSV_2=$1

git diff --color-words="[^[:space:],]+" ${CSV_1} ${CSV_2}

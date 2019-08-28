#!/bin/sh

# awk
awk 'BEGIN {
} {
    for (i = 1; i <= NF; ++i) {
        if (s[i] == "") {
            s[i] = $i
        } else {
            s[i] = s[i]" "$i;
        }
    }
} END {
    for (i = 1; s[i] != ""; i++) {
        print s[i];
    }
}' file.txt

# read
# while read -a line; do
#     for ((i = 0; i < "${#line[@]}"; ++i)); do
#         a[$i]="${a[$i]} ${line[$i]}"
#     done
# done < file.txt
# for ((i = 0; i < ${#a[@]}; ++i)); do
#     echo ${a[i]}
# done

#!/bin/bash

convertsecs() {
    h=$(bc <<< "${1}/3600")
    m=$(bc <<< "(${1}%3600)/60")
    s=$(bc <<< "${1}%60")
    printf "%02d:%02d:%02d\n" $h $m $s
}

cat ${1} | awk '{print $1"\t"$3}'
while read p
do
    echo $p
done < ${1}
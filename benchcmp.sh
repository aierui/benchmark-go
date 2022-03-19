#!/bin/bash


# benchcmp
# install it to run: go install golang.org/x/tools/...@latest

old=$1
new=$2

for i in `ls ./ | grep 'test.go'`; do
    fileName=${i%_test.go};
    echo "${fileName}"
    benchcmp ./data/${fileName}_{$old,$new}.txt
    echo ""
done;


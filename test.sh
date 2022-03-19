#!/bin/bash

DATA=data

# rm -rf $DATA
mkdir -p $DATA

version=`go version | awk '{print $3}'`

echo $version

for i in `ls ./ | grep test.go`; do 
    fileName=${i%_test.go};
    go test $i -bench=. -run=^$ > $DATA/${fileName}_${version}.txt
done;


#!/bin/bash
echo "starting $1 workflows..."
for i in `seq 1 $1`;
do
    go run start/main.go
    echo "workflow $i started"
done
echo "done"

#!/usr/bin/env bash


for file in $(ls bench); do
    hyperfine --warmup 3 "cat $file | ./peek" "cat $file | tee /dev/stderr"
done

#!/bin/bash

day=$1
echo "Setting up day $day"

dir_name=day$day
mkdir $dir_name

curl "https://adventofcode.com/2021/day/$day/input" -H "cookie: _ga=GA1.2.497526682.1638300006; session=53616c7465645f5f61815ac63165e9195fd68a37edb118b2efa0bf852667d8f825e241f4d091cd375f41cc2c30cc8dc9; _gid=GA1.2.1970541942.1638757191" > $dir_name/input.txt

cp template-dir/sol.go $dir_name
cp template-dir/sol_test.go $dir_name
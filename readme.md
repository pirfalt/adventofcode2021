# Advent of code 2021

## Fetch input

Setup.

```sh
echo 'session=<your-session>' > session.sh
```

Daily run.

```sh
source session.sh

day='15'
dir=$(printf 'day%02d' $day)
mkdir $dir
cp -r day00/ $dir/
curl --cookie "session=$session" --compressed  "https://adventofcode.com/2021/day/$day/input" > "$dir/input.txt"
```

## Work on a day

```sh
cd $dir
code *

# Test
go test
watchexec -c 'go test'

# Run
go run ./main.go
watchexec -c 'go run ./main.go'
```

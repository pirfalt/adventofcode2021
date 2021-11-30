# Advent of code 2021

## Fetch input

Setup.

```sh
echo 'session=<your-session>' > session.sh
```

Daily run.

```sh
source session.sh

day='1'
dir=$(printf 'day%02d' $day)
curl --cookie "session=$session" --compressed  "https://adventofcode.com/2020/day/$day/input" > "$dir/input.txt"
```

## Work on a day

```sh
cd $dir

# Test
go test
watchexec -c 'go test'

# Run
go run ./main.go
watchexec -c 'go run ./main.go'
```

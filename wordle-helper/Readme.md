# Wordle Helper

Let’s say you are given the 5-letter word corpus (use this one from none other than Donald Knuth himself - https://www-cs-faculty.stanford.edu/~knuth/sgb-words.txt)

Write a program (let’s call it wordle-helper) that takes 3 inputs (grays, yellows, greens) and outputs all the potential words that meet the criteria.

# Example Usage

```zsh
# build program
go build .

# execute program
# ./wordle-helper <grays> <yellows> <greens>
./wordle-helper adieupsychtr nk o3
```
### Outputs:
known
knoll
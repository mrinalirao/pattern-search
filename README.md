# pattern-search

pattern-search finds all the occurrences of a particular set of characters in a string. 
It matches the subtext against the textToSearch, outputting the positions of the beginning of each match for the subtext within the textToSearch.

To improve the algorithm further, we could use the KMP Pattern Searching algorithm whose time complexity if O(n)

# Getting Started

## Running locally

Requirements
* [Go ^1.11](https://blog.golang.org/go1.11)

To run locally
```
$ go run main.go
```

###Input
Enter the text to search, press the return/enter key once you have completed the sentence, you will be prompted to enter the subtext
```
text to search: Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!
subtext:Peter

```

###Output
```
[1 20 75]
```

## Test

To run the tests:
```
$ go test
```


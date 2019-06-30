package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to return absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Function which matches a subtext within a textToSearch and returns the  an array of positions of the beginning of the match
func PatternSearch(textToSearch string, subtext string) []int {
	var positions []int
	text := []rune(textToSearch)
	pattern := []rune(subtext)
	lenText := len(textToSearch)
	lenPattern := len(subtext)

	if lenPattern > lenText {
		return nil
	}

	for i := 0; i < lenText-lenPattern; i++ {
		k := i
		match := true
		for j := 0; j < lenPattern; j++ {
			if (text[k] == pattern[j]) || Abs(int(text[k]-pattern[j])) == 32 {
				k += 1
			} else {
				match = false
				break
			}
		}
		if match {
			pos := k - lenPattern
			positions = append(positions, pos+1)
		}
	}
	return positions
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("text to search:")
	textToSearch, _ := reader.ReadString('\n')
	textToSearch = strings.Replace(textToSearch, "\n", "", -1)
	fmt.Printf("subtext:")
	subtext, _ := reader.ReadString('\n')
	subtext = strings.Replace(subtext, "\n", "", -1)
	fmt.Println(PatternSearch(textToSearch, subtext))
}

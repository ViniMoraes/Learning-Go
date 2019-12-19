package main

import (
	"fmt"
	"strings"
	"os"
)

func main()  {
	words := os.Args[1:]

	wordsStats := collectStats(words)

	fmt.Println("Couting first letters from words...")
	for letter, count := range wordsStats {
		fmt.Printf("%s = %d\n", letter, count)
	}
}

func collectStats(words []string) map[string]int  {
	stats := make(map[string]int)
	for _, word := range words {
		firstLetter := strings.ToUpper(string(word[0]))
		count, found := stats[firstLetter]

		if found {
			stats[firstLetter] = count + 1
		} else {
			stats[firstLetter] = 1
		}
	}
	return stats
}
package main

import (
	"fmt"
	"strings"
)

type pair struct {
	word  string
	count int
}

type counter map[string]int

// -----------------------------------------------------------------------

func wordGenerator(phrase string) func() string {
	words := strings.Fields(phrase)
	index := 0
	return func() string {
		if index >= len(words) {
			return ""
		}

		word := words[index]
		index++

		return word
	}
}

func countDigits(word string) int {
	count := 0
	for _, char := range word {
		if char >= '0' && char <= '9' {
			count++
		}
	}
	return count
}

// -----------------------------------------------------------------------

func fillStats(in chan pair) counter {
	stats := make(counter)
	for {
		p := <-in
		if p.word == "" {
			break
		}

		stats[p.word] = p.count
	}
	return stats
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("Word: %s, Digits: %d\n", word, count)
	}
}

// -----------------------------------------------------------------------

func submitWords(next func() string) chan string {
	out := make(chan string)
	go func() {
		for {
			word := next()
			out <- word
			if word == "" {
				break
			}
		}
	}()
	return out
}

func countWords(in chan string) chan pair {
	out := make(chan pair)
	go func() {
		for {
			word := <-in
			count := countDigits(word)
			out <- pair{word, count}
			if word == "" {
				break
			}
		}
	}()
	return out
}

func countDigitsInWords(next func() string) counter {
	pending := submitWords(next)
	counted := countWords(pending)
	return fillStats(counted)
}

// -----------------------------------------------------------------------

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}

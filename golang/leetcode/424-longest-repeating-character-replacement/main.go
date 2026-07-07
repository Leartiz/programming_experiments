package main

import "fmt"

// https://leetcode.com/problems/longest-repeating-character-replacement/

func canGetSameChars(freqs [26]int, k int) bool {
	maxCharIndex := -1
	maxFreq := 0

	for i := 0; i < len(freqs); i += 1 {
		if freqs[i] > maxFreq {
			maxCharIndex = i
			maxFreq = freqs[i]
		}
	}

	for i := 0; i < len(freqs); i += 1 {
		if i == maxCharIndex {
			continue
		}

		k -= freqs[i] // если символ 0, то "игнор"
	}

	if maxCharIndex == -1 {
		return false
	}

	if k < 0 {
		return false
	}

	return true
}

func characterReplacement(s string, k int) int {
	/*
	   Hash Table
	   String
	   Sliding Window

	   s consists of only uppercase English letters.
	*/

	/*
	   по умолчанию считаем, что из всей строки можно получить
	   длинную последовательность одинаковых символов.

	   ...

	   ACBABBA

	   ACB
	   map = {
	       A = 1
	       C = 1
	       B = 1
	   }

	   найти тут самый максимальный
	   например A, остальные занулить попробовать
	*/

	left := 0
	right := 0

	maxLen := 1
	frequencies := [26]int{}

	for right < len(s) {
		frequencies[s[right]-'A'] += 1
		if canGetSameChars(frequencies, k) {
			maxLen = max(right-left+1, maxLen)
			right += 1
			continue
		}

		frequencies[s[left]-'A'] -= 1
		frequencies[s[right]-'A'] -= 1
		left += 1
	}

	return maxLen
}

func main() {
	/*
		Input: s = "ABAB", k = 2
		Output: 4

		Input: s = "AABABBA", k = 1
		Output: 4
	*/
	{
		fmt.Println(characterReplacement("ABAB", 2))
		fmt.Println(characterReplacement("AABABBA", 1))
	}
}

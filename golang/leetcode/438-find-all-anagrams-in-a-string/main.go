package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	// Constraints:
	//   s and p consist of lowercase English letters.
	//   1 <= s.length, p.length <= 3 * 10^4
	//

	// NOTE:
	//   важно только наличие букв, порядок нет
	//   можно проверять совпадение частот, например
	//   размер подстроки в s есть размер p
	//
	if len(p) > len(s) { // из s нельзя получить p
		return []int{}
	}

	pFreqs := map[byte]int{} // по сути массив, размера 26
	for i := 0; i < len(p); i++ {
		pFreqs[p[i]] += 1
	}

	// текущие частоты в строке s по длине p
	curFreqs := map[byte]int{}
	for j := 0; j < len(p); j++ {
		curFreqs[s[j]] += 1
	}

	result := []int{} // стартовые позиции в строке s
	if reflect.DeepEqual(pFreqs, curFreqs) {
		result = append(result, 0)
	}

	for i := 1; i <= len(s)-len(p); i++ {

		curFreqs[s[i+len(p)-1]] += 1
		curFreqs[s[i-1]] -= 1
		if curFreqs[s[i-1]] == 0 {
			delete(curFreqs, s[i-1])
		}

		if reflect.DeepEqual(pFreqs, curFreqs) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	{
		s := "cbaebabacd"
		p := "abc"
		fmt.Println(findAnagrams(s, p)) // [0 6]
	}
	{
		s := "abab"
		p := "ab"
		fmt.Println(findAnagrams(s, p)) // [0 1 2]
	}
}

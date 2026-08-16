package main

import "fmt"

// https://leetcode.com/problems/backspace-string-compare/

func increaseHashCount(s string, idx int, hashCount int) (int, int) {
	for idx >= 0 && s[idx] == '#' {
		hashCount += 1
		idx -= 1
	}
	return idx, hashCount
}

func backspaceToBegin(s string, idx int, hashCount int) int {
	for idx >= 0 {
		if s[idx] != '#' && hashCount > 0 {
			hashCount -= 1
			idx -= 1
		} else if s[idx] == '#' {
			hashCount += 1
			idx -= 1
		} else {
			break
		}
	}
	return idx
}

func backspaceCompare(s string, t string) bool {
	ptrS := len(s) - 1
	ptrT := len(t) - 1

	hashCountS := 0
	hashCountT := 0

	for ptrS >= 0 && ptrT >= 0 {
		if s[ptrS] != '#' && t[ptrT] != '#' {
			if hashCountS > 0 {
				hashCountS -= 1
				ptrS -= 1
				continue
			}
			if hashCountT > 0 {
				hashCountT -= 1
				ptrT -= 1
				continue
			}

			if s[ptrS] != t[ptrT] {
				return false
			}

			ptrS -= 1
			ptrT -= 1
			continue
		}

		// ***

		/*
			for s[ptrS] == '#' {
				hashCountS += 1
				ptrS -= 1
			}
		*/
		ptrS, hashCountS = increaseHashCount(s, ptrS, hashCountS)
		ptrT, hashCountT = increaseHashCount(t, ptrT, hashCountT)
	}

	ptrS = backspaceToBegin(s, ptrS, hashCountS)
	ptrT = backspaceToBegin(t, ptrT, hashCountT)

	if ptrS != ptrT {
		return false
	}

	return true
}

func main() {
	/*
		Input: s = "ab#c", t = "ad#c"
		Output: true
		Explanation: Both become "ac".

		Input: s = "ab##", t = "c#d#"
		Output: true
		Explanation: Both become "".

		Input: s = "a#c", t = "b"
		Output: false
		Explanation: s becomes "c" while t becomes "b".
	*/
	{
		fmt.Println(backspaceCompare("ab##", "c#d#")) // true
		fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
		fmt.Println(backspaceCompare("a#c", "b"))     // false
	}
}

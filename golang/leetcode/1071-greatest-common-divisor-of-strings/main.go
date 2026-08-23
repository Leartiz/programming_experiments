package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/greatest-common-divisor-of-strings/

func modStr(str1 string, str2 string) string {
	if len(str2) == 0 {
		panic("mod by empty")
	}

	countFound := 0
	for {
		res := strings.Index(str1[len(str2)*countFound:], str2)
		if res != 0 {
			break
		}

		countFound += 1
	}

	return str1[len(str2)*countFound:]
}

func gcdOfStrings(str1 string, str2 string) string {
	// Constraints:
	//   1 <= str1.length, str2.length <= 1000
	//   str1 and str2 consist of English uppercase letters.
	//

	// str1 будет всегда длиннее чем str2
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	for len(str2) != 0 {
		// a, b = b, a%b
		// По сути просто нужна операция % на строках
		newStr2 := modStr(str1, str2)
		if len(newStr2) >= len(str2) {
			return ""
		}
		str1, str2 = str2, newStr2
	}
	return str1
}

func main() {
	// НОД(a, b) = НОД(b, a % b)
	/*
		Пример для строк:
		НОД(ABCABC, ABC) = НОД(ABC, ABCABC % ABC)

		---

		Input: str1 = "ABABAB", str2 = "ABAB"
		НОД(ABABAB, ABAB) = НОД(ABABAB, ABABAB % ABAB)
			ABABAB % ABAB -> AB

		НОД(ABAB, AB) = НОД(ABAB, ABAB % AB)
			ABAB % AB -> ""

		Учитывать также важно и количество вхождений?

		---

		Input: str1 = "LEET", str2 = "CODE"
		НОД(LEET, CODE) = НОД(LEET, LEET % CODE) и вообще ничего не поменяется!
		s = t + t + t + ... + t + t
		Вхождений НОЛЬ.
		И тут нужен ответ. Как пустая строка.
	*/

	/*
		Input: str1 = "ABCABC", str2 = "ABC"
		Output: "ABC"

		Input: str1 = "ABABAB", str2 = "ABAB"
		Output: "AB"

		Input: str1 = "LEET", str2 = "CODE"
		Output: ""
	*/
	{
		fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // ABC
		fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // AB
		fmt.Println(gcdOfStrings("LEET", "CODE"))   // ""
	}
}

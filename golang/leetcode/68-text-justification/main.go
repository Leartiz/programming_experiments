package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/text-justification/

// хотим получить количество слов,
// которые вмещаются в строку, с учетом единичного пробела между ними
func nextLineWords(words []string, wordIndex int, maxWidth int) ([]string, int) {
	lineWords := []string{}
	curLineLen := 0
	for wordIndex < len(words) {
		curLineLen += len(words[wordIndex])

		// по сути минимальное количество пробелов = количество уже лежащих слов
		// (т.к. след. только обрабатываем)
		if curLineLen+len(lineWords) > maxWidth {
			break
		}

		lineWords = append(lineWords, words[wordIndex])
		wordIndex += 1
	}

	return lineWords, wordIndex
}

func lineJustify(words []string, maxWidth int) string {
	allWordsLen := 0 // можно получить из вне (например из nextLineWords)
	for i := 0; i < len(words); i += 1 {
		allWordsLen += len(words[i])
	}

	spaceBetween := maxWidth - allWordsLen
	spaceRest := 0
	if len(words) > 1 {
		spaceRest = spaceBetween % (len(words) - 1)
		spaceBetween = spaceBetween / (len(words) - 1)
	}

	var sb strings.Builder
	for i := 0; i < len(words); i += 1 {
		sb.WriteString(words[i])

		// если единственное слово, то пробелы нужны записать в конец.
		// иначе, после последнего слова, пробелы уже не нужны!
		if i > 0 && i == len(words)-1 {
			continue
		}

		for j := 0; j < spaceBetween; j += 1 {
			sb.WriteByte(' ')
		}

		if spaceRest > 0 {
			sb.WriteByte(' ')
			spaceRest -= 1
		}
	}

	return sb.String()
}

func lastLineJustify(words []string, maxWidth int) string {
	var sb strings.Builder
	curWidth := 0
	for i := 0; i < len(words); i += 1 {
		sb.WriteString(words[i])
		curWidth += len(words[i]) // не помню, можно ли получить из билдера?

		if i == len(words)-1 {
			continue
		}

		sb.WriteByte(' ')
		curWidth += 1
	}

	for curWidth < maxWidth {
		sb.WriteByte(' ')
		curWidth += 1
	}

	return sb.String()
}

func fullJustify(words []string, maxWidth int) []string {
	// NOTE:
	//   A word is defined as a character sequence consisting of non-space characters only.
	//   Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
	//   The input array words contains at least one word.
	//

	// EXAMPLE:
	//   "This", "is", "an", "example", "of", "text", "justification.", mxWidth = 16
	//
	//   next line - функция может вернуть например слова,
	//   которые можно впихнуть в maxWidth, с учетом одиночного пробела.
	//
	//   "This is an example" = больше 16
	//   "This is an" = 10 символов, 3 слова, 2 промежутка. 16 - 10, 6 / 2 = 3 пробела между.
	//   "This   is   an"
	//
	//   "example of text" = 15 символов, добавить еще одно слово нельзя!
	//   "example  of text" = 16 - 13 = 3 пробела, 3 - 1 (в конце обязательно), 2 / 1 = 2
	//   3 / 2 (промежутка) = 1, 3 % 2 = 1 (пойдет в первый).
	//   5 / 3 = 1; 5 % 2 = 2, счетчки, куда ставиться дополнительный пробел, или номер промежутка.\
	//
	//  "justification.  "
	//
	//  хм

	var wordIndex int
	var lineWords []string
	result := []string{}
	for {
		lineWords, wordIndex = nextLineWords(words, wordIndex, maxWidth)
		if wordIndex >= len(words) {
			result = append(result, lastLineJustify(lineWords, maxWidth))
			break
		}

		result = append(result, lineJustify(lineWords, maxWidth))
	}
	return result
}

func main() {
	{
		words := []string{"This", "is", "an", "example", "of", "text", "justification."}
		fmt.Println(fullJustify(words, 16))
	}
	{
		words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
		fmt.Println(fullJustify(words, 16))
	}
}

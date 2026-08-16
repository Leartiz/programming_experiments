package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Explore_StringsIndex(t *testing.T) {
	s := "ABCABCA"
	substr := "ABC"

	idx1 := strings.Index(s, substr)
	fmt.Println(idx1)

	fmt.Printf("len(substr): %v\n", len(substr))
	fmt.Printf("s[len(substr):]: %v\n", s[len(substr):])

	// Тоже 0, по идее если 0, то искать можно дальше,
	// иначе это и есть остаток.
	idx2 := strings.Index(s[len(substr):], substr)
	fmt.Println(idx2)

	idx3 := strings.Index(s[len(substr)*2:], substr)
	fmt.Println(idx3) // -1

	fmt.Println(s[len(substr)*2:]) // Остаток!
}

func Test_modStr(t *testing.T) {
	type data struct {
		str1 string
		str2 string
	}

	dataList := []data{
		{str1: "ABCABCA", str2: "ABC"},
		{str1: "ABCD", str2: "ABC"},
		{str1: "LEET", str2: "CODE"},
	}

	for i := 0; i < len(dataList); i += 1 {
		res := modStr(dataList[i].str1, dataList[i].str2)
		fmt.Println(res)
	}
}

func Test_gcdOfStrings(t *testing.T) {
	type data struct {
		str1 string
		str2 string
		want string
	}

	dataList := []data{
		{str1: "ABCABC", str2: "ABC", want: "ABC"},
		{str1: "ABABAB", str2: "ABAB", want: "AB"},
		{str1: "ABCD", str2: "ABC", want: ""},
		{str1: "LEET", str2: "CODE", want: ""},
	}

	for i := 0; i < len(dataList); i += 1 {
		got := gcdOfStrings(dataList[i].str1, dataList[i].str2)
		if dataList[i].want != got {
			t.Errorf("want: %v, got: %v",
				dataList[i].want, got)
		}
	}
}

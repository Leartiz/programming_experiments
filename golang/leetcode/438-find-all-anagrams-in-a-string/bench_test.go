package main

import (
	"reflect"
	"testing"
)

// NOTE:
/*
	без этого оптимизатор может понять:
		результат никуда не идёт -> весь цикл можно убрать
*/
var benchSink bool

func freqMap(s string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}

func freqArray(s string) [26]int {
	var a [26]int
	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
	}
	return a
}

func freqSlice(s string) []int {
	a := make([]int, 26)
	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
	}
	return a
}

// ручной цикл для сравнения массива/слайс
// -----------------------------------------------------------------------

func arraysEqualLoop(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func slicesEqualLoop(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// equal
// -----------------------------------------------------------------------

func BenchmarkCompare_MapDeepEqual_Equal(b *testing.B) {
	a := freqMap("abc")
	c := freqMap("abc")
	for b.Loop() {
		benchSink = reflect.DeepEqual(a, c)
	}
}

func BenchmarkCompare_ArrayEqual_Equal(b *testing.B) {
	a := freqArray("abc")
	c := freqArray("abc")
	for b.Loop() {
		benchSink = a == c
	}
}

func BenchmarkCompare_ArrayLoop_Equal(b *testing.B) {
	a := freqArray("abc")
	c := freqArray("abc")
	for b.Loop() {
		benchSink = arraysEqualLoop(a, c)
	}
}

func BenchmarkCompare_SliceLoop_Equal(b *testing.B) {
	a := freqSlice("abc")
	c := freqSlice("abc")
	for b.Loop() {
		benchSink = slicesEqualLoop(a, c)
	}
}

// unequal
// -----------------------------------------------------------------------

func BenchmarkCompare_MapDeepEqual_Unequal(b *testing.B) {
	a := freqMap("abc")
	c := freqMap("abd")
	for b.Loop() {
		benchSink = reflect.DeepEqual(a, c)
	}
}

func BenchmarkCompare_ArrayEqual_Unequal(b *testing.B) {
	a := freqArray("abc")
	c := freqArray("abd")
	for b.Loop() {
		benchSink = a == c
	}
}

func BenchmarkCompare_ArrayLoop_Unequal(b *testing.B) {
	a := freqArray("abc")
	c := freqArray("abd")
	for b.Loop() {
		benchSink = arraysEqualLoop(a, c)
	}
}

func BenchmarkCompare_SliceLoop_Unequal(b *testing.B) {
	a := freqSlice("abc")
	c := freqSlice("abd")
	for b.Loop() {
		benchSink = slicesEqualLoop(a, c)
	}
}

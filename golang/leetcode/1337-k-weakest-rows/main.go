package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Example 1:
/*
	Input: mat =
		[[1,1,0,0,0],
		[1,1,1,1,0],
		[1,0,0,0,0],
		[1,1,0,0,0],
		[1,1,1,1,1]],
	k = 3
	Output: [2,0,3]

	Explanation:
		The number of soldiers in each row is:
			- Row 0: 2
			- Row 1: 4
			- Row 2: 1
			- Row 3: 2
			- Row 4: 5
	The rows ordered from weakest to strongest are [2,0,3,1,4].
*/

// Sort indices
// -----------------------------------------------------------------------

func kWeakestRows(mat [][]int, k int) []int {
	// NOTE:
	/*
		Constraints:
			m == mat.length
			n == mat[i].length
			2 <= n, m <= 100
			mat[i][j] is either 0 or 1.

		A row i is weaker than a row j if one of the following is true:
			The number of soldiers in row i is less than the number of soldiers in row j.
			Both rows have the same number of soldiers and i < j.
	*/

	// количество солдат в каждой строке
	soldiers := make([]int, len(mat))
	for i, row := range mat {
		soldierCount := 0
		for _, kind := range row {
			if kind == 1 {
				soldierCount += 1
			}
		}

		soldiers[i] = soldierCount
	}

	// номера строк, по сути как объявлены
	weakestRows := make([]int, len(mat)) // не capacity
	for i := 0; i < len(weakestRows); i += 1 {
		weakestRows[i] = i
	}

	sort.Slice(weakestRows, func(i, j int) bool {
		a, b := weakestRows[i], weakestRows[j]

		// Both rows have the same number of soldiers and i < j.
		//
		if soldiers[a] == soldiers[b] {
			return a < b
		}

		// The number of soldiers in row i is less than the number of soldiers in row j.
		//
		return soldiers[a] < soldiers[b]
	})

	return weakestRows[:k]
}

// Min-heap
// -----------------------------------------------------------------------

type RowStrength struct {
	soldiersCount int
	rowIndex      int
}

type MatHeap []RowStrength

func (h MatHeap) Len() int { return len(h) }
func (h MatHeap) Less(i, j int) bool {
	if h[i].soldiersCount == h[j].soldiersCount {
		return h[i].rowIndex < h[j].rowIndex
	}

	return h[i].soldiersCount < h[j].soldiersCount
}
func (h MatHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MatHeap) Push(x any) {
	hElement := x.(RowStrength)
	*h = append(*h, hElement)
}

func (h *MatHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kWeakestRowsMinHeap(mat [][]int, k int) []int {
	h := &MatHeap{}
	heap.Init(h)

	for i, row := range mat {
		soldiersCount := 0
		for _, kind := range row {
			if kind == 1 {
				soldiersCount += 1
			}
		}

		heap.Push(h, RowStrength{
			soldiersCount: soldiersCount,
			rowIndex:      i,
		})
	}

	weakestRows := make([]int, k)
	for i := 0; i < len(weakestRows); i += 1 {
		hElement := heap.Pop(h) // return any
		weakestRows[i] = hElement.(RowStrength).rowIndex
	}

	return weakestRows
}

// Quick select/partial sort (think)
// -----------------------------------------------------------------------

const debugQuickSelect = true

func formatRowStrength(rs RowStrength) string {
	return fmt.Sprintf("(%d,%d)", rs.soldiersCount, rs.rowIndex)
}

func printRows(label string, rows []RowStrength, l, r int) {
	if !debugQuickSelect {
		return
	}
	fmt.Printf("--- %s  [l=%d r=%d] ---\n", label, l, r)
	for i := l; i <= r; i++ {
		fmt.Printf("  [%d] %s\n", i, formatRowStrength(rows[i]))
	}
}

func weaker(a, b RowStrength) bool {
	if a.soldiersCount != b.soldiersCount {
		return a.soldiersCount < b.soldiersCount
	}
	return a.rowIndex < b.rowIndex
}

func partition(rows []RowStrength, l, r int) int {
	pivotIdx := (l + r) / 2
	if debugQuickSelect {
		fmt.Printf("\npartition: pivotIdx=%d → %s, swap with [%d]\n", pivotIdx, formatRowStrength(rows[pivotIdx]), r)
	}
	rows[pivotIdx], rows[r] = rows[r], rows[pivotIdx]

	store := l
	for i := l; i < r; i++ {
		if !weaker(rows[r], rows[i]) { // rows[i] не сильнее pivot (rows[r])
			rows[store], rows[i] = rows[i], rows[store]
			store++
		}
	}
	rows[store], rows[r] = rows[r], rows[store]
	if debugQuickSelect {
		fmt.Printf("partition: pivot on p=%d\n", store)
		printRows("after partition", rows, l, r)
	}
	return store
}

// NOTE:
/*
	Не сортировать всё, а разрезать массив так,
	чтобы k слабых оказались слева; потом досортировать только их.
*/
func kWeakestRowsQuickSelect(mat [][]int, k int) []int {
	rows := make([]RowStrength, len(mat))
	for i, row := range mat {
		soldiersCount := 0
		for _, kind := range row {
			if kind == 1 {
				soldiersCount += 1
			}
		}

		rows[i] = RowStrength{
			soldiersCount: soldiersCount,
			rowIndex:      i,
		}
	}

	if debugQuickSelect {
		fmt.Println("=== quickselect start ===")
		printRows("1) собрали rows", rows, 0, len(rows)-1)
		fmt.Printf("k=%d, ищем top-k в индексах [0..%d]\n", k, k-1)
	}

	endIndex := len(rows) - 1
	l, r := 0, endIndex
	for l <= r {
		if debugQuickSelect {
			fmt.Printf("\n2) quickselect loop: l=%d r=%d\n", l, r)
		}
		p := partition(rows, l, r)
		if debugQuickSelect {
			fmt.Printf("   p=%d, k-1=%d → ", p, k-1)
		}
		if p == k-1 {
			if debugQuickSelect {
				fmt.Println("p == k-1, top-k готов в rows[0:k]")
			}
			break
		}
		if p > k-1 {
			if debugQuickSelect {
				fmt.Printf("p > k-1 → ищем слева, r=%d\n", p-1)
			}
			r = p - 1
		} else {
			if debugQuickSelect {
				fmt.Printf("p < k-1 → ищем справа, l=%d\n", p+1)
			}
			l = p + 1
		}
	}

	if debugQuickSelect {
		fmt.Println("\n3) top-k до sort (порядок не гарантирован):")
		for i := 0; i < k; i++ {
			fmt.Printf("   [%d] %s\n", i, formatRowStrength(rows[i]))
		}
	}

	sort.Slice(rows[:k], func(i, j int) bool {
		return weaker(rows[i], rows[j])
	})

	if debugQuickSelect {
		fmt.Println("4) top-k после sort (weakest → strongest):")
		for i := 0; i < k; i++ {
			fmt.Printf("   [%d] %s → row %d\n", i, formatRowStrength(rows[i]), rows[i].rowIndex)
		}
		fmt.Println("=== answer ===")
	}

	weakestRows := make([]int, k)
	for i := range weakestRows {
		weakestRows[i] = rows[i].rowIndex
	}
	return weakestRows
}

// -----------------------------------------------------------------------

func main() {
	{
		mat := [][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		}
		fmt.Println(kWeakestRowsQuickSelect(mat, 3)) // Output: [2,0,3]
	}
}

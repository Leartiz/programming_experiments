package main

import "fmt"

// https://leetcode.com/problems/shift-2d-grid/

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%v ", grid[i][j])
		}
		fmt.Println()
	}
}

func shiftGrid(grid [][]int, k int) [][]int {
	// NOTE:
	/*
	   Constraints:
	       1 <= m <= 50
	       1 <= n <= 50

	       m == grid.length
	       n == grid[i].length

	       -1000 <= grid[i][j] <= 1000
	       0 <= k <= 100
	*/

	rows := len(grid)    // количество строк
	cols := len(grid[0]) // количество элементов в строке
	numberCount := rows * cols

	// NOTE:
	//   оптимизация, если количество элементов в сетке == k,
	//   то ничего не делать
	//
	if k%numberCount == 0 { // 0 сюда же
		return grid
	}

	// NOTE: example 1
	/*
	   m = rows = 4
	   n = cols = 4
	   gridSize = m * n = 16

	   [3   8  1    9]
	   [19  7  2    5]
	   [4   6  11  10]
	   [12  0  21  13]

	   ---
	   k = 4
	   shiftRows =  k / cols (количество элементов в строке) = 4 / 4 = 1
	   shiftCols = k % cols = 0 - не двигаем колонки
	   ---
	   k = 5
	   shiftRows =  k / cols (количество элементов в строке) = 5 / 4 = 1 - количество строк для сдвига
	   shiftRows = shiftRows % rows = 1 % 4 = 1
	   shiftCols = k % cols = 1
	   ---
	   k = 27
	   shiftRows =  k / cols (количество элементов в строке) = 27 / 4 = 6 - количество строк для сдвига
	   shiftRows = shiftRows % rows = 6 % 4 = 2
	   shiftCols = k % cols = 3
	*/

	// NOTE: example 2
	/*
	   m = rows = 3
	   n = cols = 3
	   gridSize = m * n = 9

	   [1,2,3]
	   [4,5,6]
	   [7,8,9]

	   k = 1

	   shiftRows =  k / cols = 1 / 3 = 0
	   shiftRows = shiftRows % 3 = 0 % 3 = 0
	   shiftCols = 1 % cols = 1
	*/

	shiftRows := k / cols
	shiftRows %= rows
	shiftCols := k % cols

	for shiftRows > 0 {
		// один цикл, один сдвиг
		temp := grid[rows-1] // запоминаем последний (станет первый)
		for i := 0; i < rows; i++ {
			temp, grid[i] = grid[i], temp // swap
		}
		shiftRows -= 1
	}

	for shiftCols > 0 {
		temp := grid[rows-1][cols-1]
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				temp, grid[i][j] = grid[i][j], temp
			}
		}
		shiftCols -= 1
	}

	return grid
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printGrid(grid)

	grid = shiftGrid(grid, 1)
	printGrid(grid)
}

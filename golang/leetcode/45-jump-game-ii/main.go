package main

import "fmt"

// https://leetcode.com/problems/jump-game-ii/

func jumpPickNext(nums []int) int {
	// NOTE:
	/*
	   1 <= nums.length <= 10^4
	   0 <= nums[i] <= 1000
	   It's guaranteed that you can reach nums[n - 1]
	*/

	// NOTE:
	/*
	   по идее, со старта, надо прыгать туда, где есть максимальное число
	   например если стоит на 2, а дальше 4 и 3, можно прыгнуть на 4
	   хотя...

	   4 + 1 = 5
	   3 + 2 = 5

	   2 4 3 1 1 1

	   ex. 1
	   ---

	   2 3 1 1 4

	   2 -> 3 или 2 -> 1, надо прыгать в 3, потому что смогу потом дальше прыгнуть
	   3 -> 1, 1, 4

	   1. с текущей позиции, посмотреть куда можно прыгнть (если сразу в конец, то вернуть количество прыжков)
	   2. выбрать наиболее лучшую, следующую позицию по алгоритму max(nums[i]+i, ...)
	   3. перейти на выбранню позицию
	   4. вернуться к шагу 1
	*/

	jumpCount := 0 // результат
	curPos := 0    // текущая позиция, где мы

	for {
		maxJumpSize := nums[curPos] // максимальный размер прыжка с текущей позиции

		if curPos >= len(nums)-1 { // в целом, если мы уже прыгнули сюда, то это выход
			break
		}
		if curPos+maxJumpSize >= len(nums)-1 { // в целом, если мы уже прыгнули сюда, то это выход
			jumpCount += 1
			break
		}

		// ищем следующую позицию
		nextPos := curPos + 1
		nextMaxJumpPos := nextPos + nums[nextPos] // "следующее" максимальное место куда можно прыгнуть

		curJumpBorder := curPos + maxJumpSize // текущая граница прыжка

		for i := curPos + 1; i <= curJumpBorder; i += 1 {
			if nums[i]+i > nextMaxJumpPos {
				nextMaxJumpPos = i + nums[i]
				nextPos = i
			}
		}

		curPos = nextPos
		jumpCount += 1
	}
	return jumpCount
}

func jumpReach(nums []int) int {
	n := len(nums)

	jumps := 0  // количество прыжков, которые мы уже сделали
	curEnd := 0 // граница текущего прыжка

	curFarthest := 0 // максимальный индекс, который можно достичь из всех позиций

	for i := 0; i < n-1; i++ { // нужно достичь последнего индекса n-1.
		curFarthest = max(curFarthest, i+nums[i]) // максимальное значение, куда можем прыгнуть
		if i == curEnd {                          // сработает и на нуле

			curEnd = curFarthest // и потом мы пойдем по всех числам, будем искать куду дальше сможет прыгнуть
			jumps++

			if curEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

// -----------------------------------------------------------------------

func main() {
	{
		nums := []int{2, 3, 1, 1, 4}
		fmt.Println(jumpPickNext(nums))
	}
	{
		nums := []int{1, 2, 1, 1, 1}
		fmt.Println(jumpPickNext(nums))
	}
	{
		nums := []int{2, 1}
		fmt.Println(jumpPickNext(nums))
	}
}

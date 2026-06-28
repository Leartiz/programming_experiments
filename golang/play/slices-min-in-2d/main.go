package main

func main() {
	data := [][]int{
		{1, 2, 3},
		{1, 2, 3, 5},
		//...
	}
	var resultFound bool = false
	var result int

	for i := range data {
		for j := range data[i] {
			if !resultFound {
				result = data[i][j]
				resultFound = true
			}

			if result > data[i][j] {
				result = data[i][j]
			}
		}
	}
}

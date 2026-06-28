package main

import (
	"fmt"
)

func main() {
	{
		type ID int

		var n int = 10
		var id ID = 10

		// Так нельзя:
		// id = n
		// Compile-time error:
		// cannot use n (variable of type int) as ID value in assignment

		// А так можно:
		id = ID(n)
		fmt.Printf("id is %T\n", id)
		// id is main.ID
	}

	{
		// ID - псевдоним для int
		type ID = int

		var n int = 10
		var id ID = 10

		id = n // так можно, ведь тип один и тот же
		fmt.Printf("id is %T\n", id)
		// id is int
	}
}

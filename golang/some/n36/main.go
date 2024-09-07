package main

import (
	"fmt"
)

type greeter interface {
	greet()
}

// -----------------------------------------------------------------------

type english struct {
	name string
}

func (e *english) greet() {
	if e == nil {
		fmt.Println("(english) self eq nil")
		return
	}
	fmt.Println("Hello", e.name)
}

// -----------------------------------------------------------------------

type russian struct {
	name string
}

func (e *russian) greet() {
	if e == nil {
		fmt.Println("(russian) self eq nil")
		return
	}
	fmt.Println("Привет", e.name)
}

// -----------------------------------------------------------------------

func main() {
	{
		var ivar any
		// type == nil, value == nil
		// поэтому ivar == nil
		fmt.Println(ivar == nil)
		// true

		var e *english
		fmt.Println(e == nil)
		// true

		// ***

		// Но как только интерфейсной переменной присвоили значение,
		//     type перестает быть nil.
		// Поэтому переменная больше не равна nil,
		//     даже если value равно nil.

		ivar = e
		// type == *english, value == nil
		// поскольку type != nil, то ivar != nil
		fmt.Println(ivar == nil) // !
		// false
	}
	{
		//var ivar greeter
		//ivar.greet() // panic!
	}
	{
		var ivar greeter
		ivar = (*english)(nil)
		ivar.greet()

		ivar = (*russian)(nil)
		ivar.greet()

		// ***

		ivar = &english{name: "Joshua"}
		ivar.greet()

		ivar = &russian{name: "Иван"}
		ivar.greet()

		ivar = nil
		//ivar.greet() // panic!
	}
}

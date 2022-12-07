package main

import "fmt"

func main() {
	// a pointer is a value that contains the address of a variable

	var i int8 = 54
	// the & operator yields the address of a variable,
	p := &i
	fmt.Println("memory address of i is:", p)
	// the * operator retrieves the variable that thepointer refers to
	fmt.Println("p points to a variable with value:", *p)
	*p = 100
	fmt.Println("i is:", i)

	// p points to a specific data type, here int8, and cannot hold a value from another data type
	// *p = "Go" //string
	// *p = 3456787654567 //int64

	// a pointer can be nil
	var p1 *int8 = nil
	fmt.Println(p1)
	// var i1 int8 = nil //illegal
}

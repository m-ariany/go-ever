// source go101.org

package main // specify the source file's package

import (
	"math/rand"
) // import a standard package

const MaxRnd = 16 // a named constant declaration

// "main" function is the entry function of a program.
func main() {

	var num = 100
	x, y := statRandomNumbers(num)
	// Call two built-in functions (print and println).
	print("Result: ", x, " + ", y, " = ", num, "? ")
	println(x+y == num)
}

// A function declaration
/*
 statRandomNumbers produces a certain number of
 non-negative random integers which are less than
 MaxRnd, then counts and returns the numbers of
 small and large ones among the produced randoms.
 n specifies how many randoms to be produced.
*/
func statRandomNumbers(n int) (int, int) {

	// Declare two variables (both as 0).
	var a, b int
	// A for-loop control flow.
	for i := 0; i < n; i++ {
		// An if-else control flow.
		if rand.Intn(MaxRnd) < MaxRnd/2 {
			a = a + 1
		} else {
			b++ // same as: b = b + 1
		}
	}
	return a, b // this function return two results
}

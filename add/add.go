// Package add provides the method that helps perform
// an operation of addition for two values.
package add

import "golang.org/x/exp/constraints"

// Add adds the two given values and return the value of
// their sum.
//
// More information on addition can be found at [addition]
//
// [addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

// Integer is a constraint that permits any integer or float type.
type Number interface {
	constraints.Integer | constraints.Float
}
package goadd

import goconstraint "golang.org/x/exp/constraints"

type Number interface {
  goconstraint.Integer | goconstraint.Float
}

// Add takes two integers, a and b, and returns their sum.
// 
// Parameters:
//   - a: The first integer.
//   - b: The second integer.
//
// Returns:
//   The sum of a and b.
func Add[T Number](a, b T) T {
  return a + b
}

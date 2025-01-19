package goadd

import"golang.org/x/exp/constraints"

type Number interface {
  constraints.Integer | constraints.Float
}

// Add takes two numbers, a and b, and returns their sum.
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

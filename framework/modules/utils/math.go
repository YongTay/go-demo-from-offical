package utils

func Add[T int | float32 | float64](n ...T) T {
	var sum T
	for _, v := range n {
		sum = v + sum
	}
	return sum
}

package adder

func Add(x, y int) int {
	// Check code coverage
	if x == y {
		return x * 2
	} else {
		return x + y
	}
}

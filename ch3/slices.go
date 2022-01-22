package main

import "fmt"

// func main() {
// 	var x = [...]int{1, 10: 100, 6, 5: 4, 15}
// 	var y = [2][12]int{{1: 12}, x}
// 	y[0][0] = 11
// 	fmt.Println(y[0][y[0][0]])
// 	var n []int
// 	fmt.Println(n != nil, len(n), cap(n), n)
// 	n = append(n, 1, 2, 3)
// 	fmt.Println(n)
// 	var nonnilslice = make([]int, 0, 10)
// 	fmt.Println(nonnilslice != nil, len(nonnilslice), cap(nonnilslice), nonnilslice)
// 	fmt.Println([]int{1, 2, 3, 4}[:2])
// }

// func main() {
// 	x := []int{1, 2, 3, 4}
// 	y := x[:2]
// 	z := x[1:]
// 	x[1] = 20
// 	y[0] = 10
// 	z[1] = 30
// 	fmt.Println("x:", x)
// 	fmt.Println("y:", y)
// 	fmt.Println("z:", z)
// }

// func main() {
// 	x := make([]int, 0, 5)
// 	x = append(x, 1, 2, 3, 4)
// 	y := x[:2]
// 	z := x[2:]
// 	fmt.Println(cap(x), cap(y), cap(z))
// 	y = append(y, 30, 40, 50)
// 	x = append(x, 60)
// 	z = append(z, 70)
// 	fmt.Println("x:", x)
// 	fmt.Println("y:", y)
// 	fmt.Println("z:", z)
// }

// y = [1, 2, 30, 40, 50] => [1, 2, 30, 40, 60] => => [1, 2, 30, 40, 70]
// x = [1, 2, 30, 40, 60] => [1, 2, 30, 40, 70]
// z = [30, 40, 70]

// func main() {
// 	x := make([]int, 0, 5)
// 	x = append(x, 1, 2, 3, 4)
// 	y := x[:2:2]
// 	z := x[2:4:5]
// 	fmt.Println(cap(x), cap(y), cap(z))
// 	y = append(y, 30, 40, 50)
// 	x = append(x, 60)
// 	z = append(z, 70)
// 	fmt.Println("x:", x)
// 	fmt.Println("y:", y)
// 	fmt.Println("z:", z)
// }

// func main() {
// 	x := [...]int{5, 6, 7, 8}
// 	y := x[:2]
// 	z := x[2:]
// 	y[0] = 0
// 	y = append(y, 1)
// 	fmt.Println("x:", x)
// 	fmt.Println("y:", y)
// 	fmt.Println("z:", z)
// }

func main() {
	x := []int{1, 2, 3, 4}
	num := copy(x[1:3], x[:3])
	fmt.Println(x, num)
}

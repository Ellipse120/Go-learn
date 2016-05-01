package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 2
	y = x / 2
	return
}

func Cycle() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt2(x float64) float64 {
	// initialize variables
	z, d := float64(1), float64(1)

	// if delta is greater than 1e-15, seek closer approximation
	for d > 1e-15 {
		z0 := z
		z = z - (z * z - x) / (2 * z)
		d = math.Abs(z - z0)
	}
	return z
}

var java, python = 1, false

func main() {
	fmt.Println("My favorite number is", math.E)
	fmt.Println(add(1, 2))
	fmt.Println(swap("a", "b"))

	a, b := swap("hello", "go")
	fmt.Println(a, b)

	fmt.Println(split(16))
	i, j := .1, true
	fmt.Println(i, j, java, python)

	fmt.Println(Cycle())

	fmt.Println(sqrt(2), sqrt(-16))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// 牛顿求根法
	fmt.Println(Sqrt2(2))
	fmt.Println(math.Sqrt(2))

}




package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe bool = true
	MaxInt uint64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5+12i)

	i int
	s string ="a"
)


func main() {
	const f = "%T(%v)\n"
	fmt.Println(f,ToBe,ToBe)
	fmt.Println(f,MaxInt,MaxInt)
	fmt.Println(f,z,z)
	fmt.Println("%v %q\n",i,s)

	x,y  := 3,4
	g := math.Sqrt(float64(x*x + y*y))
	z := uint(g)
	fmt.Println(x,y, z)

}

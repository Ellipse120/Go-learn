package main

import (
	"time"
	"fmt"
)

type Vertex struct {
	X,Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	p = &Vertex{1,4}
)

func main() {
	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() <17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	//
	//fmt.Println("defer test: ")
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("done defer!!!")
	//
	//fmt.Println(Vertex{1,2})
	v:=Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println(v1,p,v2,v3)

}

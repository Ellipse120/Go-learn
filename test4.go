package main

import (
	"time"
	"fmt"
	"strings"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	p = &Vertex{1, 4}
	pow = []int{1,2,4,8,16,32,64,128}
)

func main() {
	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
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
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println(v1, p, v2, v3)

	//数组
	var a1 [2]string
	a1[0] = "hello"
	a1[1] = "world"
	fmt.Println(a1[0], a1[1])
	fmt.Println(a1)

	// slice
	s := []int{1, 2, 3, 6, 11, 112}
	fmt.Println("s ==", s)

	/*for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n",i,s[i])
	}*/
	// Slice 切片
	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[:3] ==", s[:3])
	fmt.Println("s[4:] ==", s[4:])

	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)

	// printSlice
	/*
	* slice 由函数 make 创建。
	* 这会分配一个全是零值的数组并且返回一个 slice 指向这个数组.
	* 为了指定容量，可传递第三个参数到 make：
	*/
	a := make([]int,5)
	printSlice("a", a)
	b := make([]int,0,5)
	printSlice("b",b)
	c :=b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d",d)
	// append 对slice 添加元素
	d = append(d,2,35,4)
	printSlice("d",d)
	// range 可以通过赋值给 _ 来忽略序号和值。
	for i, v := range pow  {
		fmt.Printf("2**%d = %d\n",i,v)
	}
	//如果只需要索引值，去掉 “ , value ” 的部分即可。
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow{
		fmt.Printf("%d\n",value)
	}
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s , len(x),cap(x),x)
}

/*import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	canvas := make([][]uint8, dx)
	for i:=0; i<dx; i++ {
		canvas[i] = make([]uint8, dy)
		for j:=0; j<dy; j++ {
			canvas[i][j] = uint8(i^j)
		}
	}
	return canvas
}

func main() {
	pic.Show(Pic)
}*/
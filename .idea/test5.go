package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
	40.68433, -74.39967,
	},
	"Google": Vertex{
	37.42202, -122.08408,
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

// 闭包
func fibonacci() func() int {
	f0, f1, f2 := 0, 1, 0
	index := 0
	return func() int{
		if index == 0 {
			index += 1
			return f0
		} else if index == 1 {
			index += 1

			return f1
		} else {
			f2 = f0 + f1
			f0 = f1
			f1 = f2
			return f2
		}
	}
}

func main() {
	//m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)

	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5,12))
	fmt.Println(hypot)
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// 闭包
	f:= fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// map练习 wordcount
/*import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	s_arr := strings.Fields(s)//分割字符串为字符数组
	s_map := make(map[string]int)//建立map
	//对s_arr中的每个字符进行循环
	for i:= 0; i<len(s_arr); i++ {
		if s_map[s_arr[i]] == 0 { //当还没有统计过该字符时，赋值为1
			s_map[s_arr[i]] = 1
		} else {                  //当统计过该字符时，更新计数值+1
			s_map[s_arr[i]] = s_map[s_arr[i]] + 1
		}
	}
	return s_map
}

func main() {
	s := "I love my work and I"
	res := WordCount(s)
	fmt.Println(res)
}*/
package main

import (
    "fmt"
    "golang.org/x/tour/wc"
    "strings"
)

type Vertex struct{
  Lat, Long float64
}

var m map[string]Vertex  //映射将键映射到值。映射的零值为 nil 。nil 映射既没有键，也不能添加键。
var m1 = map[string]Vertex{ 
    "Bell Labs": Vertex{40.68433, -74.39967,}, //映射的文法与结构体相似，不过必须有键名。
    "Google": Vertex{37.42202, -122.08408,},
    "Google2": {37.42202, -122.08408,},	//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
}

func WordCount(s string) map[string]int{
	fmt.Println(strings.Fields(s))

	var m = make(map[string]int)

	for i, v := range strings.Fields(s) {
		fmt.Println(i, v)
		m[v] = m[v] + 1
	}
	return m
    //return map[string]int{"x": 1}
}

func main(){
    m = make(map[string]Vertex)	//make 函数会返回给定类型的映射，并将其初始化备用。
    m["Bell Labs"] = Vertex{40.68433, -74.39967,}
    fmt.Println(m["Bell Labs"])
    fmt.Println(m1)

    //mutating-maps 修改映射
    m2 := make(map[string]int)
    m2["Answer"] =42
    fmt.Println("The value:", m2["Answer"])
    v2, ok2 := m2["Answer"]  //通过双赋值检测某个键是否存在：若key在m中ok2为true否则ok2为false。
    fmt.Println("The value:", v2, "Present?", ok2)
    m2["Answer"] =48
    fmt.Println("The value:", m2["Answer"])
    delete(m2, "Answer")
    fmt.Println("The value:", m2["Answer"])
    v, ok := m["Answer"]  //通过双赋值检测某个键是否存在：若key在m中ok为true否则ok为false。
    fmt.Println("The value:", v, "Present?", ok) //若key不在映射中，那么v 是该映射元素类型的零值。 同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
    v1, ok1 := m["xxx"]
    fmt.Println("The value:", v1, "Present?", ok1) //当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。


    //exercise-maps.go
    wc.Test(WordCount)
}

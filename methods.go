//方法: Go 没有类。不过你可以为结构体类型定义方法。
//方法就是一类带特殊的 接收者 参数的函数。
//方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

//在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
func (v Vertex) Abs() float64 {		//为Vertex定义一个Abs 方法
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//方法即函数 记住：方法只是个带接收者参数的函数。 
//等价的写法 现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
func Absf(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//非结构方法
type MyFloat float64
func (f MyFloat) Abs() float64 {		//为Vertex定义一个Abs 方法
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//methods-pointers.go
func (v *Vertex) Scale(f float64) {		//为Vertex定义一个Abs 方法
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v1 := Vertex{3,4}	//实例化一个
	fmt.Println(v1.Abs())	//调用方法
	fmt.Println(Absf(v1))	//调用函数
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v1.Scale(10)
	fmt.Println(v1)
	fmt.Println(v1.Abs())
}


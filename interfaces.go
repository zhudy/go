//接口类型 是由一组方法签名定义的集合。 接口类型的变量可以保存任何实现了这些方法的值。

package main

import (
    "fmt"
    "math"
)

type Abser interface {
	Abs() float64
}


//类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
type I interface {
  M()
}
type T struct {
  S string
}
//隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。 因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
//func (t T)M(){ // 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t *T)M(){
  if t == nil {
	fmt.Println("<nil>")
	return
  }
  fmt.Println(t.S)
}

//接口也是值。它们可以像其它值一样传递。 接口值可以用作函数的参数或返回值。
// 在内部，接口值可以看做包含值和具体类型的元组： (value, type)
//接口值保存了一个具体底层类型的具体值。 接口值调用方法时会执行其底层类型的同名方法。
type F float64
func (f F) M(){
  fmt.Println(f)
}
//func describe(i I){ //david--
func describe(i interface{}){
  fmt.Printf("(%v, %T)\n", i, i)
}

func main(){
    var a Abser //接口类型的变量
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3,4}

    a = f //接口类型的变量a就像一个指针？ 可以调用任何实现了这些方法的值。
    fmt.Println(a.Abs())
    a = &v
    fmt.Println(a.Abs())
    //a = v //由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser。
    fmt.Println(a.Abs())
  
/*
  var i I = T{"hello"}
  describe(i)
  i.M() 
*/
  var i I
  var t *T
  i = t
  describe(i)
  i.M() 
  i = &T{"hello"}
  describe(i)
  i.M() 

fmt.Println("接口值")
  var i1 I = &T{"hello"}
  describe(i1)
  i1.M() 
  var j F = F(math.Pi)
  describe(j)
  j.M() 

/*
//nil 接口值将报错
  var k I
  describe(k)
  k.M() //panic: runtime error: invalid memory address or nil pointer dereference [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x109ee7f]
*/

//空接口
  //type X interface {}
  var abcd interface{} 
  describe(abcd)
  abcd = 42
  describe(abcd)
  abcd = "hello"
  describe(abcd)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f<0 {
	return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

package main

import (
    "fmt"
)

func main(){
    var i interface{} = "hello"

    s:= i.(string)//该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 s
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64) //为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。 若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。 否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生panic。
    fmt.Println(f, ok)

/*
    f1 := i.(float64) //若i并未保存T类型的值，该语句就会触发一个panic
    fmt.Println(f1)
*/

    //类型选择
    do(21)
    do("hello")
    do(true)
}

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

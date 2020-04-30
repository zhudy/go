package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	n:=0
	n1:=1
	
    return func() int { 
        result := n
        n = n1
        n1 = result + n1
        return result
    }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


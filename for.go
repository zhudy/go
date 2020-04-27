package main

import "fmt"

func main(){
    sum := 0
    for i := 0; i < 10; i++ {
	sum += i
        fmt.Println(i, sum)
    }

    for ; sum < 10000; {
	sum += sum
        fmt.Println(sum)
    }

    sum = 1
    for sum <100 { 	// == while
	sum += sum
	fmt.Println(sum)
    }

    for {	//无限循环
	fmt.Println(sum)
	break
	fmt.Println(sum)
    }
}

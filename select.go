package main

import (
    "fmt"
    "time"
)

func fibonacci(c, quit chan int){
    x, y := 0, 1
    for {
        select {  //阻塞到某分支可以执行为止，多个分支都准备好时随机选择一个执行
	case c <- x:
	    x, y = y, x+y
        case <-quit:
	    fmt.Println("quit")
	    return
        }
    }
}

func main(){
    c := make(chan int)
    quit := make(chan int)
    go func(){
        for i:=0; i<10; i++{
	    fmt.Println(<-c)
	}
	quit <-0
    }()
    fibonacci(c, quit)

    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond) //此函数马上返回，返回一个time.Time类型的Chan，不阻塞。
    for {
	select {
	case <-tick:
	    fmt.Println("tick.")
	case <-boom:
	    fmt.Println("BOOM!")
	    return
	default: //为了在尝试发送或者接收时不发生阻塞，可使用 default 分支
	    fmt.Println("    .")
	    time.Sleep(50* time.Millisecond)
	}
    }
}

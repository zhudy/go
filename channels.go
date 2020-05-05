//信道是带有类型的管道，用信道操作符 <- 来发送或者接收值。
// 和映射与切片一样，信道在使用前必须创建： ch := make(chan int)
//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
package main

import (
    "fmt"
)

func sum(s []int, c chan int) {
    sum :=0
    for _, v := range s {
	sum += v
    }
    c <- sum  //将和送入c
}

func fibonacci(n int, c chan int){
    x, y := 0, 1
    for i:=0; i<n; i++{
	c <- x
	x, y = y, x+y
    }
    //发送者可通过 close 关闭一个信道来表示没有需要发送的值了
    //只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发panic
    //信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
    close(c)	
}

func main(){
    s:= []int{7,2,8,-9, 4,0}
    c:= make(chan int)
    go sum(s[:len(s)/2],c)
    go sum(s[len(s)/2:],c)
    x, y:= <-c, <-c //从c中接收
    fmt.Println(x, y, x+y)

    //将长度作为第二个参数来初始化一个带缓冲的信道ch := make(chan int, 100)
    //仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    ch <- 3
    fmt.Println(<-ch)
    fmt.Println("hello hahaha")

    //range and close
    c1 := make(chan int, 10)
    go fibonacci(cap(c1), c1)
    for i := range c1 {  //不断从信道接收值，直到它被关闭。
	fmt.Println(i)
    }
}

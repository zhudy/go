package main

import (
    "fmt"
    "time" 
)

func say(s string) {
    for i:=0; i<5; i++{
	time.Sleep(100* time.Millisecond)
    	fmt.Println(s)
    }
}

func main(){
    go say("world") //启动一个新的 Go 程并执行
    say("hello")
}

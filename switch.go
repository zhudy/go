package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		//freebasd, openbsd, plan9, windows...
		fmt.Printf("%s. \n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("time.Saturday Type: %T Value: %v\n", time.Saturday, time.Saturday)
	fmt.Printf("today Type: %T Value: %v\n", today, today+0)
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow.")
	case today+2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}

	t := time.Now()
	switch true{	//true 可以不写,没有条件的 switch 同 switch true 一样。
	case t.Hour() <12:
		fmt.Println("Good morning!")
	case t.Hour() <17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}


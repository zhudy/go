package main

import (
    "fmt"
    "strings"
)

func main(){
    var aa [2]string
    aa[0] = "Hello"
    aa[1] = "World"
    fmt.Println(aa[0], aa[1])
    fmt.Println(aa)

    primes := [6]int{2,3,5,7,11,13}
    fmt.Println(primes)

    var s []int = primes[1:4]
    fmt.Println(s)

    names := [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
    }
    fmt.Println(names)

    //slice
    a := names[0:2]
    b := names[1:3]
    fmt.Println(a,b)

    b[0] = "XXX"
    fmt.Println(a,b)
    fmt.Println(names)

    //slice-literals
    q := []int{2,3,5,7,11,13}
    fmt.Println(q)
    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    ss := []struct {
	i int
	b bool
    }{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
    }
    fmt.Println(ss)

    //slice bounds
    fmt.Println(q)
    q = q[1:4]
    fmt.Println(q)
    q = q[:2]
    fmt.Println(q)
    q = q[1:]
    fmt.Println(q)

    //slice len cap
    //切片拥有 长度 和 容量。 切片的长度就是它所包含的元素个数。 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。 切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
    q = []int{2,3,5,7,11,13}
    printSlice(q)
    q = q[:0]
    printSlice(q)
    q = q[:4]
    printSlice(q)
    q = q[2:]
    printSlice(q)

    //nil
    var e []int
    fmt.Println(e, len(e), cap(e))
    if e == nil {
	fmt.Println("nil!")
    }

    //makeing-slices
    s1 := make([]int, 5)    
    printSslice("s1", s1)

    s2 := make([]int, 0, 5)    
    printSslice("s2", s2)

    s3 := s2[:2]
    printSslice("s3", s3)

    s4 := s3[2:5]
    printSslice("s4", s4)

    //slices of slice
    board := [][]string{
	[]string{"_", "_", "_"}, 
	[]string{"_", "_", "_"}, 
	[]string{"_", "_", "_"}, 
	[]string{"_", "_", "_"}, 
    }
    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"
    for i := 0; i < len(board); i++{
	fmt.Printf("%s\n", strings.Join(board[i], " "))
    }

    //append
    var as []int
    printSlice(as)
    as = append(as, 0)
    printSlice(as)
    as = append(as, 1)
    printSlice(as)
    as = append(as, 2, 3, 4)
    printSlice(as)

    //range
    var pow = []int{1,2,4,8,16,32,64,128}
    for i, v := range pow {
	fmt.Printf("2**%d = %d\n", i, v)
    }

    //range-continued
    pow = make([]int ,10)
    for i := range pow {
	pow[i] = 1 << uint(i)  // == 2**i
    }
    for _, value := range pow {
	fmt.Printf("%d\n", value)
    }
   
}

func printSlice(s []int){
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSslice(ss string, s []int){
    fmt.Printf("%s len=%d cap=%d %v\n", ss, len(s), cap(s), s)
}

package main

import (
    "fmt"
)

type Person struct{
  Name string
  Age int
}
//{Arthur Dent 42} {Zaphd Beeblebrox 9001}
// Stringer 是一个可以用字符串描述自己的类型。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age) // Arthur Dent (42 years) Zaphd Beeblebrox (9001 years)
}

type IPAddr [4]byte
//如果没定义下面这个函数，则默认返回： googleDNS: [8 8 8 8] loopback: [127 0 0 1]
func (ip IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", ip[0],ip[1],ip[2],ip[3]) 
}

func main(){
    a:= Person{"Arthur Dent", 42}
    z:= Person{"Zaphd Beeblebrox", 9001}
    fmt.Println(a, z)

    hosts := map[string]IPAddr{
	"loopback": {127,0,0,1},
	"googleDNS": {8,8,8,8},
    }
    for name, ip := range hosts{
	fmt.Printf("%v, %v\n", name, ip)
    }
}

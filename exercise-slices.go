package main

import (
    "golang.org/x/tour/pic" 
    //"math"
)

//https://blog.csdn.net/vinson0526/article/details/52279076
func Pic(dx, dy int) [][]uint8 {
    var pic [][]uint8
    for i:=0; i<dy; i++{
	var one_line []uint8
	for j:=0; j<dx; j++{
	    one_line = append(one_line, (uint8)(i & j)) //(i+j)/2, i*j, i^j, i*log(j), i%(j+1)

	    //i*log(j)
	    //one_line = append(one_line, (uint8)(i*  int (math.Log2(float64(j)))))
	}
	pic = append(pic, one_line)
    }
    return pic
}

func main(){
    pic.Show(Pic)
}

package main

import (
    "fmt"
    "time"
    "math"
)

type MyError struct{
  When time.Time
  What string
}
//与 fmt.Stringer 类似，error 类型是一个内建接口
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{time.Now(), "it didn't work",}
}

//
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x<0 {
	    return 0, ErrNegativeSqrt(x)
	}
        z :=1.0
        oldz := z
        fmt.Println(z, z*z, x-z*z)

        for {
                oldz = z
                z -= (z*z -x)/(2*z)
                if oldz == z {
                        break
                }
                fmt.Println(z, z*z, x-z*z)

                if math.Abs(x - z*z) < 0.000000000000001 {
                        break
                }
        }
        return z, nil
}


func main(){
    if err := run(); err != nil {
	fmt.Println(err)
    }

    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))

}

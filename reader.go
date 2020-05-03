package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "golang.org/x/tour/reader"
)

type MyReader struct{}
func (r MyReader) Read(p []byte) (int, error) {
    for i := 0; i< len(p); i++{
	p[i] = 'A'
    }
    return len(p), nil
}

type rot13Reader struct{
    r io.Reader
}
func (rs rot13Reader) Read(buf []byte) (int, error) {
    length, err := rs.r.Read(buf)
    if err != nil {
        return length, err
    }

    for i := 0; i < length; i++ {
        v := buf[i]
        switch {
        case 'a' <= v && v <= 'm':
            fallthrough
        case 'A' <= v && v <= 'M':
            buf[i] = v + 13
        case 'n' <= v && v <= 'z':
            fallthrough
        case 'N' <= v && v <= 'Z':
            buf[i] = v - 13
        }
    }
    return length, nil
}

func main(){
    r := strings.NewReader("Hello, Reader!")
    b := make ([]byte, 8)

    for{ 
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
	if err == io.EOF{
		break
	}
    }
    reader.Validate(MyReader{})

    rs := strings.NewReader("Lbh penpxrq gur pbqr!")
    rr := rot13Reader{rs}
    io.Copy(os.Stdout, &rr)
    fmt.Println("\n")
}

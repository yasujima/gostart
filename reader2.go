package main

import "golang.org/x/tour/reader"
import "fmt"
import "io"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}


func main() {
	reader.Validate(MyReader{})
	//	r := MyReader{}
	r := MyReader{}
	b := make([]byte, 10)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}

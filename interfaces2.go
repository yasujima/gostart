package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
	fmt.Printf("%T\n", t)
}

func main() {
	var i I = T{"hello"}
	i.M();
}

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string

}

func (t *T) M() {
	if t == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	i.M()

	i = &T{"hellO"}
	i.M()
}

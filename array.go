package main

import "fmt"


var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// func main() {
// 	var a [2]string
// 	a[0] = "hello"
// 	a[1] = "world"
// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)

// 	s := []struct {
// 		i int
// 		b bool
// 	}{
// 		{2, true},
// 		{3, false},
// 		{5, true},
// 		{7, true},
// 		{11, false},
// 		{13, true},
// 	}
// 	fmt.Println(s)
// 	fmt.Println(s[6])
// }
	

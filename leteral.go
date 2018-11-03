package main

import "fmt"
import "reflect"

func main() {
	var v map[string]string
	v = map[string]string{"11":"111"}
	fmt.Println(len(v))


	var a [4]int
	fmt.Println(a)


	b := [2]string{"pen", "teller"}
	c := [...]string{"xx", "yyy"}
	fmt.Println(b,c)

	af := f()
	fmt.Println(reflect.ValueOf(af).Kind())
	fmt.Println(af)

	letters := []string{"a","b","c"}
	fmt.Println(letters)
	s := make([]byte, 5, 5)
	fmt.Println(s)
	fmt.Println(reflect.ValueOf(s).Kind())
	ss := []byte{0,1,2,3,4}
	fmt.Println(ss)
	ss = ss[:len(ss)]
	fmt.Println(ss)
	sss := ss[:]
	ss[0] = 100
	fmt.Println(sss)

	function()
	hoge()
	
}

func f() []int {

	var a = []int{1,2,3}
	return a

}

func function() {

	var x = struct{name string}{"name"}
	fmt.Println(x)
//	[...]int{}{}
//	[]int{}{}
	//	map[string]string{}

	var as = [...][]int{[]int{1,2}, []int{2,3}}
	fmt.Println(as)

	v := struct {
		name string
		age int
	} {
		name: "hoge",
		age:10,
	}
	fmt.Println(v)

	v2 := map[interface{}]interface{}{"age":12}
	fmt.Println(v2)

	v3 := &struct{}{}
	fmt.Println(v3)

	func(i int) int { return i} (0)
}

type Fuga interface {
	Fuga()
}

type FugaFuga struct {
	name string
}
func (f *FugaFuga) Fuga() {
	fmt.Println(f.name)
	f.name = "called fuga change"
}

func fff(f Fuga) {
	f.Fuga()
}

func hoge() {

	f := FugaFuga{"my name is fugafuga"}
	fff(&f)
	fff(&f)
	
}
		

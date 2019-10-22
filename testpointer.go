package main


type Person struct {
	name string
	age int64
}

func main() {
}

func NewPersonVV() Person {
	return Person{name: "hoge", age:12}
}

func NewPersonPV() Person {
	//return &Person{name: "hoge", age:12}
	return nil
}

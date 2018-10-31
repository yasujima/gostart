package main
import (
	"encoding/json"
	"log"
	"fmt"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Person struct {
	ID int `json:"id"`
	Name string
	Email string `json:"-"`
	Age int
	Address string `json:"address,omitempty"`
	Memo string
}

func main() {
	person := &Person{
		ID:1,		Name: "Gop",
		Email: "gop@example.com",
		Age:5,
		Address: "aaa",
		Memo: "gogo",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	var man Person

	bb := []byte(`{"id":1,"name":"GONAME","age":5}`)
	e := json.Unmarshal(bb, &man)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(man)


	file, e := os.Create("./person.json")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
	
// 	message := []byte("hello world\n")
// 	l,e := file.Write(message)
// //	fmt.Println(l)
// 	log.Fatal(l)
// 	if e != nil {
// 		log.Fatal(e)
// 	}
}

		
		

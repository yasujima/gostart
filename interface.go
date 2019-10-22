package main

import (
	"log"

)


type Inf interface {
	exec()
}

type Impl struct {
}

// func (impl *Impl) exec() {
// 	log.Printf("* impl exec called")
// }

func (impl Impl) exec() {
	log.Printf("impl exec called")
}

func returnInterface() (Inf) {
	return &Impl{}
}

func returnStruct() Impl {
	return Impl{}
	//return &Impl{}  こっちはNG
}


func main() {
	var m Inf = &Impl{}
	m.exec()

	var i Inf = returnInterface()
	i.exec()

	var j Inf = returnStruct()
	j.exec()

}

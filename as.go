//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import "fmt"

func anyExec(any interface{}) {
	fmt.Print("xxxx\n")
}

func main() {
	anyExec(12)
	anyExec("hello")
}

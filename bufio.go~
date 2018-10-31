package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b:=bufio.NewWriter(os.Stdout)
	for i:=0; i<100; i++ {
		fmt.Fprintln(b, strings.Repeat("x",100))
	}
	b.Flush()
}

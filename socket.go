package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Listen error %s\n", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept error %s\n", err)
		return
	}

	defer conn.Close()

	fmt.Println(" message from client:")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Printf("Read error: %s.\n", err)
		}
		fmt.Print(string(buf[:n]))
	}
}






package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting pop3 server")
	conn, err := net.Listen("tcp", ":3110")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening to pop3 port: %s", err.Error())
	}
	for {
		c, _ := conn.Accept()
		go handleconnection(c)
	}
}

func handleconnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		// Reads a line from the client
		raw_line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error!!" + err.Error())
			return
		}
		fmt.Print(raw_line)
	}
}

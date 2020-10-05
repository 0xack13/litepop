package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
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
	l := log.New(os.Stdout, "", 0)
	reader := bufio.NewReader(c)
	for {
		// Reads a line from the client
		raw_line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error!!" + err.Error())
			return
		}
		Log(l, raw_line)
	}
	c.Close()
}

func Log(l *log.Logger, msg string) {
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " > ")
	l.Print(msg)
}

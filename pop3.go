package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	host = "localhost"
)
const (
	STATE_UNAUTHORIZED = 1
	STATE_TRANSACTION  = 2
	STATE_UPDATE       = 3
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

func handleconnection(c net.Conn) string {
	defer c.Close()
	// initialize state wuth UNAUTHORIZED
	// var (
	// 	eol = "\r\n"
	// 	// state = 1
	// )
	l := log.New(os.Stdout, "", 0)
	reader := bufio.NewReader(c)
	_, err := fmt.Fprintf(c, "+OK simple POP3 server %s\n", host)
	// message, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Print("->: " + message)
	for {
		rawline, err := reader.ReadString('\n')
		// line := strings.Trim(rawline, "\r")
		if err != nil {
			fmt.Println(err.Error())
			return "error"
		}
		// fmt.Println(line)
		// test uses \n
		// real uses \r\n
		if rawline == "QUIT\r\n" {
			// fmt.Fprint(c, "Good bye!")
			// fmt.Print("Good bye!")
			// c.Close()
			return "quitting"
		}
		Log(l, rawline)
	}
}

// Log function
func Log(l *log.Logger, msg string) {
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " > ")
	l.Print(msg)
}

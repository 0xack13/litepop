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
		go handleconnection(c, c)
	}
}

func handleconnection(c net.Conn, cc net.Conn) {
	// defer c.Close()
	// initialize state wuth UNAUTHORIZED
	// var (
	// 	eol = "\r\n"
	// 	// state = 1
	// )
	// l := log.New(os.Stdout, "", 0)
	go func() {
		// cc.Write([]byte("+OK Litepop server ready\r\n"))
		// cc.Close()
		fmt.Fprintf(cc, "+OK Litepop server ready\r\n")
	}()
	reader := bufio.NewReader(c)

	for {
		// go func() {
		// 	fmt.Fprint(c, "For loop!\n")
		// }()
		rawline, err := reader.ReadString('\n')
		if err != nil {
			// fmt.Println("errrrorr: " + err.Error())
			// return "error"
			continue
		} else if rawline == "QUIT\r\n" {
			go func() {
				cc.Write([]byte("Good bye!\r\n"))
				c.Close()
			}()
			// fmt.Print("Good bye!")
			// c.Close()
			// return "quitting"
		} else if rawline == "USER\r\n" {
			go func() {
				// fmt.Fprint(c, "Good bye!\n")
				cc.Write([]byte("gooood byeee\n"))
				cc.Close()
				// c.Close()
			}()
			// fmt.Print("Good bye!")
			// c.Close()
			// return "USER_OK"
		} else {
			go func() {
				// fmt.Fprint(c, "Good bye!\n")
				cc.Write([]byte("Command not supported\r\n"))
				// cc.Close()
				// c.Close()
			}()
		}
		// Log(l, rawline)
		// return "true"
	}
}

// Log function
func Log(l *log.Logger, msg string) {
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " > ")
	l.Print(msg)
}

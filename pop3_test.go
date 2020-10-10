package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"testing"
)

type recordingConn struct {
	net.Conn
	io.Writer
}

func (c *recordingConn) Write(buf []byte) (int, error) {
	return c.Writer.Write(buf)
}

func TestClient_handle_quits_succcesffuly(t *testing.T) {
	// var r ConnMock
	r, w := net.Pipe()
	var buf bytes.Buffer
	r = &recordingConn{
		Conn:   r,
		Writer: io.MultiWriter(r, &buf),
	}
	// defer w.Close()
	defer r.Close()
	// timeoutDuration := 2 * time.Second
	// r.SetReadDeadline(time.Now().Add(timeoutDuration))
	go func() {
		w.Write([]byte("QUIT\r\n"))
		// w.Close()
	}()
	reader := bufio.NewReader(r)
	//read welcome message
	msg, err := reader.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(msg)
	// msg, err = reader.ReadString('\n')
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// // print(r.Read())
	// fmt.Println(msg)
	if handleconnection(r) == "quitting" {
		fmt.Println("Success")
		return
	}
	// panic("test")
}

// func TestClient_handle_timeout(t *testing.T) {
// 	r, w := net.Pipe()
// 	defer w.Close()
// 	defer r.Close()
// 	timeoutDuration := 2 * time.Second
// 	r.SetReadDeadline(time.Now().Add(timeoutDuration))
// 	go func() {
// 		w.Write([]byte("NOVAL\n"))
// 	}()

// 	if handleconnection(r) == "quitting" {
// 		fmt.Println("Success")
// 		return
// 	}
// 	// return
// }

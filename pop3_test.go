package main

import (
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
	defer r.Close()
	var buf bytes.Buffer
	r = &recordingConn{
		Conn:   r,
		Writer: io.MultiWriter(w, &buf),
		// Writer: io.Copy,
	}
	// defer w.Close()
	// timeoutDuration := 2 * time.Second
	// r.SetReadDeadline(time.Now().Add(timeoutDuration))
	go func() {
		w.Write([]byte("QUIT\r\n"))
		// w.Close()
	}()
	// fmt.Println(handleconnection(r))
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

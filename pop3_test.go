package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClient_handle_quits_succcesffuly(t *testing.T) {
	r, w := net.Pipe()
	defer w.Close()
	defer r.Close()
	timeoutDuration := 2 * time.Second
	r.SetReadDeadline(time.Now().Add(timeoutDuration))
	go func() {
		w.Write([]byte("QUIT\n"))
	}()

	if handleconnection(r) == "quitting" {
		fmt.Println("Success")
		return
	}
	panic("test")
}

func TestClient_handle_timeout(t *testing.T) {
	r, w := net.Pipe()
	defer w.Close()
	defer r.Close()
	timeoutDuration := 2 * time.Second
	r.SetReadDeadline(time.Now().Add(timeoutDuration))
	go func() {
		w.Write([]byte("NOVAL\n"))
	}()

	if handleconnection(r) == "quitting" {
		fmt.Println("Success")
		return
	}
	return
}

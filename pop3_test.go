package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"testing"
	"time"
)

type recordingConn struct {
	net.Conn
	io.Writer
}

func (c *recordingConn) Write(buf []byte) (int, error) {
	return c.Writer.Write(buf)
}

func TestClient_handle_quits_succcesffuly(t *testing.T) {
	expected := "+OK Litepop server ready\r\nGood bye!\r\n"
	r, w := net.Pipe()
	defer r.Close()
	var buf bytes.Buffer
	r = &recordingConn{
		Conn:   r,
		Writer: io.MultiWriter(w, &buf),
	}
	go func() {
		w.Write([]byte("QUIT\r\n"))
		// w.Close()
	}()
	// fmt.Println(handleconnection(r))
	// if handleconnection(r, w) == "quitting" {
	// 	fmt.Println("Success")
	// 	return
	// }
	// t.Fail()
	go func() {
		handleconnection(r, w)
	}()
	// fmt.Println(handleconnection(r))
	// var b []byte
	b, _ := ioutil.ReadAll(r)
	// if err != nil {
	// 	fmt.Println("errrrr  ")
	// 	fmt.Println(err)
	// }
	// fmt.Println("mystring *****" + string(b))
	if string(b) != expected {
		t.Fatalf("Not matching expected quit message")
	}

}

func TestClient_handle_user_succcesffuly(t *testing.T) {
	r, w := net.Pipe()
	// defer r.Close()
	// defer w.Close()
	// var buf bytes.Buffer
	// r = &recordingConn{
	// 	Conn:   r,
	// 	Writer: io.MultiWriter(w, &buf),
	// }
	go func() {
		w.Write([]byte("USER\r\n"))
		// w.Close()
	}()
	go func() {
		handleconnection(r, w)
	}()
	// fmt.Println(handleconnection(r))
	// var b []byte
	// b, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	fmt.Println("errrrr  ")
	// 	fmt.Println(err)
	// }
	// fmt.Println("mystring *****" + string(b))
	// w.Close()
	// r.Close()
	// t.Fail()
	// buf1 := make([]byte, 0, 4096) // big buffer
	// tmp := make([]byte, 256)      // using small tmo buffer for demonstrating
	// for {
	// 	n, err := r.Read(tmp)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println("read error:", err)
	// 		}
	// 		break
	// 	}
	// 	//fmt.Println("got", n, "bytes.")
	// 	buf1 = append(buf1, tmp[:n]...)

	// }
	// fmt.Println("total size:", len(buf1))
	// fmt.Println(string(buf1))
	// go func() {
	// 	w.Write([]byte("USER\r\n"))
	// 	// w.Close()
	// }()
	// r.Close()
	// w.Close()
}

func TestClient_handle_timeout(t *testing.T) {
	r, w := net.Pipe()
	defer r.Close()
	var buf bytes.Buffer
	r = &recordingConn{
		Conn:   r,
		Writer: io.MultiWriter(w, &buf),
	}
	// timoeout for testing only
	timeoutDuration := time.Second
	r.SetReadDeadline(time.Now().Add(timeoutDuration))
	go func() {
		w.Write([]byte("NOVAL\n"))
	}()

	go func() {
		handleconnection(r, w)
	}()

	// if handleconnection(r, w) == "quitting" {
	// 	fmt.Println("Success")
	// 	return
	// }
	// t.Fail()
	expected := "+OK Litepop server ready\r\nCommand not supported\r\n"
	b, _ := ioutil.ReadAll(r)
	if string(b) != expected {
		t.Fatalf("Not matching expected quit message")
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	logReq, err := http.ReadRequest(reader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(logReq.URL.Path)
	file, err := os.Open(logReq.URL.Path[1:]) // For read access.
	if err == nil {
		defer file.Close() // make sure to close the file even if we panic.
	}

	if err != nil {
		fmt.Println(err)
		conn.Write([]byte("HTTP/1.1 404 Not Found"))
		conn.Close()
		return
	}
	stat, err := file.Stat()
	size := stat.Size()
	str := fmt.Sprintf("HTTP/1.1 200 OK\nConnection: close\nContent-Type: text/html\nContent-Length:%d\n\n", size)
	conn.Write([]byte(str))
	io.Copy(conn, file)
}

func main() {
	ln, err := net.Listen("tcp", ":6789")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 12000})
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])

		if rand.Intn(10) > 4 {
			_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
			if err != nil {
				fmt.Printf(err.Error())
			}
		}
	}
}

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 12000}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		sendTime := time.Now()
		// every round we refresh deadline time
		conn.SetReadDeadline(sendTime.Add(time.Second * 1))

		message := fmt.Sprintf("Ping %d %s", i+1, sendTime)
		conn.Write([]byte(message))
		msg := make([]byte, 1024)
		_, err := conn.Read(msg)
		if err != nil && err.(net.Error).Timeout() {
			fmt.Printf("Sequence %d: Request timed out\n", i+1)
			continue
		}
		// reader := bufio.NewReader(conn)
		// status, err := reader.ReadString(byte("\n"))
		rtt := float64(time.Since(sendTime) / time.Millisecond)
		fmt.Printf("Sequence %d: Reply from %s    RTT = %.3fms\n", i+1, "localhost:12000", rtt)
	}
}

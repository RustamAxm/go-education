package main

import (
	"log"
	"net"
	"fmt"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 10001})
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	message := "Привет, stepik!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("msg send done")
}

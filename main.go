package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	fmt.Println("hello world")
	c := &serial.Config{Name: "/dev/cu.usbmodem14301", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if string(buf[:n]) != "\r" && string(buf[:n]) != "\n" {
			log.Printf("%q", buf[:n])
		}
	}
}

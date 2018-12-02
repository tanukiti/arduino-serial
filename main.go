package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tarm/serial"
)

func main() {
	id := "5bfb97202c4437d772c936e2"
	fmt.Println("hello world")
	c := &serial.Config{Name: "/dev/cu.usbmodem1421", Baud: 9600}
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
		fmt.Println(string(buf[:n]))
		if string(buf[:n]) != "\r" && string(buf[:n]) != "\n" {
			if string(buf[:n]) == "0" {
				fmt.Println("sleep")
				url := "https://hal-iot.net/iw/nellow/sleep/" + id
				http.Get(url)
			}

			if string(buf[:n]) == "1" {
				fmt.Println("wakeup")
				url := "https://hal-iot.net/iw/nellow/wakeup/" + id
				http.Get(url)
			}
		}
	}
}

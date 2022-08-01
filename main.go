package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		machine.LED.Toggle()
		fmt.Printf("hello tinygo devcontiner\r\n")
		time.Sleep(1 * time.Second)
	}
}

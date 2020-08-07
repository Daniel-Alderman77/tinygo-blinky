package main

import (
	"machine"
	"time"
)

func main() {
	led()
}

func led() {
	led := machine.D7
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("+")
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		println("-")
		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}

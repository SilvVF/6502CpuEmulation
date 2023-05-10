package main

import (
	"6502CpuEmulation/cpu"
	"fmt"
)

func main() {
	olc := cpu.NewOlc6502()
	bus := cpu.NewBus()

	olc.ConnectBus(bus)

	fmt.Println(olc)
}

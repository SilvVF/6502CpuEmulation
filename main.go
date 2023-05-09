package main

import (
	"6502CpuEmulation/flags6502"
	"fmt"
)

func main() {
	cpu := NewOlc6502()
	bus := NewBus()

	cpu.ConnectBus(bus)
	
	fmt.Println(flags6502.B)
}

type Olc6502 struct {
	a      uint8  // Accumulator Register
	x      uint   // X Register
	y      uint8  // Y Register
	stkp   uint8  // Stack Pointer (points to location on bus)
	pc     uint16 // Program Counter
	status uint8  // Status Register
	bus    *Bus
}

func NewOlc6502() Olc6502 {
	return Olc6502{
		a:      0x00,
		x:      0x00,
		y:      0x00,
		stkp:   0x00,
		pc:     0x0000,
		status: 0x00,
	}
}

func (cpu *Olc6502) ConnectBus(b *Bus) {
	cpu.bus = b
}

func (cpu *Olc6502) Read(addr uint16, bReadOnly bool) uint8 {
	return cpu.bus.Read(addr, bReadOnly)
}

func (cpu *Olc6502) Write(addr uint16, data uint8) {
	cpu.bus.Write(addr, data)
}

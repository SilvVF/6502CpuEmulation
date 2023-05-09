package main

type Bus struct {
	cpu Olc6502
	ram [64 * 1024]uint8
}

func NewBus() *Bus {
	b := Bus{
		ram: [65536]uint8{},
	}
	for i := range b.ram {
		b.ram[i] = 0x00
	}
	return &b
}

func (b *Bus) Read(addr uint16, bReadOnly bool) uint8 {
	if addr >= 0x0000 && addr <= 0xFFFF {
		return b.ram[addr]
	}
	// read outside range return 0
	return 0x00
}

func (b *Bus) Write(addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		b.ram[addr] = data
	}
}

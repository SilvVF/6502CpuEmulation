package cpu

type Olc6502 struct {
	a      uint8  // Accumulator Register
	x      uint   // X Register
	y      uint8  // Y Register
	stkp   uint8  // Stack Pointer (points to location on bus)
	pc     uint16 // Program Counter
	status uint8  // Status Register
	bus    *Bus

	fetched uint8
	addrAbs uint16
	addrRel uint16
	opcode  uint8
}

type Instruction struct {
	name     string
	addrMode *func() uint8
	operate  *func() uint8
	cycles   uint8
}

func NewOlc6502() Olc6502 {
	return Olc6502{
		a:       0x00,
		x:       0x00,
		y:       0x00,
		stkp:    0x00,
		pc:      0x0000,
		status:  0x00,
		fetched: 0x00,
		addrAbs: 0x0000,
		addrRel: 0x00,
		opcode:  0x00,
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

func (cpu *Olc6502) Clock() {

}

func (cpu *Olc6502) Reset() {

}

func (cpu *Olc6502) Irq() {

}

func (cpu *Olc6502) Nmi() {

}

func (cpu *Olc6502) Fetch() {

}

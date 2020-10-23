package vm

const (
	// KBSR is a keyboard status
	KBSR uint16 = 0xFE00
	// KBDR is a keyboard data
	KBDR uint16 = 0xFE02
)

func (cpu *CPU) memoryWrite(address uint16, val uint16) {
	cpu.Memory[address] = val
}

func (cpu *CPU) memoryRead(address uint16) uint16 {
	if address == KBDR {
		cpu.memoryWrite(KBSR, cpu.memoryRead(KBSR)&0x7FFF)
	}
	return uint16(cpu.Memory[address])
}

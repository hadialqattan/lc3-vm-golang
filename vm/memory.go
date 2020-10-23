package vm

import (
	"github.com/pkg/term"
)

const (
	// KBSR is a keyboard status
	KBSR = 0xFE00
	// KBDR is a keyboard data
	KBDR = 0xFE02
)

func (cpu *CPU) memWrite(address uint16, val uint16) {
	cpu.Memory[address] = val
}

func (cpu *CPU) memRead(address uint16) uint16 {
	if address == KBSR {
		if checkKey() {
			cpu.Memory[KBSR] = (1 << 15)
			cpu.Memory[KBDR] = getchar()
		} else {
			cpu.Memory[KBSR] = 0
		}
	}
	return cpu.Memory[address]
}

func getchar() []byte {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)
	t.Restore()
	t.Close()
	if err != nil {
		return nil
	}
	return bytes[0:numRead]
}

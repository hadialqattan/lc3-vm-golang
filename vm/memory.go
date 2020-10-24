package vm

import (
	"encoding/binary"
	"io/ioutil"
)

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

func (cpu *CPU) loadProgramImage(path string) error {
	bin, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	origin := binary.BigEndian.Uint16(bin[:2])
	for i := 2; i < len(bin); i += 2 {
		cpu.Memory[origin] = binary.BigEndian.Uint16(bin[i : i+2])
		origin++
	}

	return nil
}

func (cpu *CPU) processKBInput() {
	valOfKBSR := cpu.memoryRead(KBSR)
	isKBSRready := (valOfKBSR & 0x8000) == 0
	if isKBSRready && len(cpu.keysBuffer) > 0 {
		cpu.memoryWrite(KBSR, valOfKBSR|0x8000)
		cpu.memoryWrite(KBDR, uint16(cpu.keysBuffer[0]))
	}
}

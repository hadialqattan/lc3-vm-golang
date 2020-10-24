package vm

import (
	"fmt"
	"time"
)

// Traps
const (
	// GETC to get character from keyboard, not echoed onto the terminal
	GETC uint16 = 0x20
	// OUT to output a character
	OUT uint16 = 0x21
	// PUTS to output a word string
	PUTS uint16 = 0x22
	// IN to get character from keyboard, echoed onto the terminal
	IN uint16 = 0x23
	// PUTSP to output a byte string
	PUTSP uint16 = 0x24
	// HALT to halt the program
	HALT uint16 = 0x25
)

func (cpu *CPU) trap(instr uint16) {
	switch instr & 0xFF {
	case GETC:
		cpu.getc()
	case OUT:
		cpu.out()
	case PUTS:
		cpu.puts()
	case IN:
		cpu.in()
	case PUTSP:
		cpu.putsp()
	case HALT:
		cpu.halt()
	}
}

//========================================

func (cpu *CPU) getc() {
	// block until a key is pressed
	for {
		if len(cpu.keysBuffer) > 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	// pop a char from the `keysBuffer` queue into register 0
	cpu.Registers[0], cpu.keysBuffer = uint16(cpu.keysBuffer[0]), cpu.keysBuffer[1:]
}

func (cpu *CPU) out() {
	char := rune(cpu.Registers[0])
	fmt.Printf("%c", char)
}

func (cpu *CPU) puts() {
	var address uint16 = cpu.Registers[0]
	var char, i uint16
	for ok := true; ok; ok = (char != 0x0) {
		char = cpu.Memory[address+i] & 0xFFFF
		fmt.Printf("%c", rune(char))
		i++
	}
}

func (cpu *CPU) in() {}

func (cpu *CPU) putsp() {}

func (cpu *CPU) halt() {
	cpu.stop()
}

package vm

import (
	"errors"
	"math"
)

// CPU is a LC-3 CPU emulator.
type CPU struct {
	Memory     [math.MaxUint16 + 1]uint16 // 65536 locations (128kb)
	Opcodes    opcodes                    // 16 instruction
	keysBuffer []rune                     // Key Buffer
	IsRunning  bool                       // Run state

	// 10 total registers.
	Registers [8]uint16 // general purpose
	PC        uint16    // program counter
	COND      uint16    // condition flags
}

// NewCPU ~> Create new LC-3 CPU.
func NewCPU() *CPU {
	cpu := &CPU{}
	// set the PC to starting position
	// 0x3000 is the default
	cpu.PC = 0x3000
	return cpu
}

func (cpu *CPU) run() error {
	if len(cpu.Memory) < 1 {
		return errors.New("TODO: NOPROG HAS LOADED")
	}

	cpu.IsRunning = true
	for cpu.IsRunning {
		cpu.processKBInput()
		cpu.executeInstruction()
		cpu.Memory[0xFFFF]++
	}

	return nil
}

func (cpu *CPU) stop() {
	cpu.IsRunning = false
}

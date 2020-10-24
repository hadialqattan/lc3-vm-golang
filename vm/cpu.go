package vm

import (
	"math"
)

// CPU is a LC-3 CPU emulator.
type CPU struct {
	Memory     [math.MaxUint16 + 1]uint16 // 65536 locations (128kb)
	Opcodes    opcodes                    // 16 instruction
	KeysBuffer []rune                     // Key Buffer
	IsRunning  bool                       // Run state

	// 10 total registers.
	Registers [8]uint16 // general purpose
	PC        uint16    // program counter
	COND      uint16    // condition flags
}

// NewCPU initialize a new LC3 CPU
func NewCPU() *CPU {
	cpu := &CPU{}
	// set the PC to starting position
	// 0x3000 is the default
	cpu.PC = 0x3000
	return cpu
}

// Run executes the loaded program
func (cpu *CPU) Run() error {
	if len(cpu.Memory) < 1 {
		return errNoProgram
	}

	cpu.IsRunning = true
	for cpu.IsRunning {
		// process any keyboard input
		cpu.processKBInput()

		// execute the current instruction
		cpu.executeInstruction()

		// increment MCC
		cpu.Memory[0xFFFF]++
	}

	return nil
}

// Stop emulates the termination signal
func (cpu *CPU) Stop() {
	cpu.IsRunning = false
}

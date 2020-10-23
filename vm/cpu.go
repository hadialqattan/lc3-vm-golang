package vm

import (
	"math"
)

// CPU is a LC-3 CPU emulator.
type CPU struct {
	Memory  [math.MaxUint16]uint16 // 65536 locations (128kb)
	Opcodes opcodes                // 16 instruction

	// 10 total registers.
	registers [8]uint16 // general purpose
	PC        uint16    // program counter
	COND      condReg   // condition flags
}

type registers struct {
	// general purpose registers
	R0 uint16
	R1 uint16
	R2 uint16
	R3 uint16
	R4 uint16
	R5 uint16
	R6 uint16
	R7 uint16

	PC   uint16  // program counter
	COND condReg // condition flags
}

type opcodes struct {
	BR   uint8 // 0  ~> branch
	ADD  uint8 // 1  ~> add
	LD   uint8 // 2  ~> load
	ST   uint8 // 3  ~> store
	JSR  uint8 // 4  ~> jump register
	AND  uint8 // 5  ~> bitwise and
	LDR  uint8 // 6  ~> load register
	STR  uint8 // 7  ~> store register
	RTI  uint8 // 8  ~> unused
	NOT  uint8 // 9  ~> bitwise not
	LDI  uint8 // 10 ~> load indirect
	STI  uint8 // 11 ~> store indirect
	JMP  uint8 // 12 ~> jump
	RES  uint8 // 13 ~> reserved (unused)
	LEA  uint8 // 14 ~> load effective address
	TRAP uint8 // 15 ~> execute trap
}

type condReg struct {
	P bool // Sign (S), negative result
	Z bool // Zero (Z), zero result
	N bool // Parity (P), the number of 1 bits is even in the result
}

// NewCPU ~> Create new LC-3 CPU.
func NewCPU() *CPU {
	cpu := &CPU{}
	// set the PC to starting position
	// 0x3000 is the default
	cpu.PC = 0x3000
	return cpu
}

func (cpu *CPU) run() {
	running := true
	for running {

	}
}

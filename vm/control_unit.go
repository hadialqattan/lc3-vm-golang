package vm

import (
	"log"
)

const (
	// FLGPOS for 'Positive' result
	FLGPOS uint16 = 1 << 0
	// FLGZRO for 'Zero' result
	FLGZRO uint16 = 1 << 1
	// FLGNEG for 'Negative' result
	FLGNEG uint16 = 1 << 2
)

// Operation codes
const (
	// BR 0  ~> branch
	BR uint16 = iota
	// ADD 1  ~> add
	ADD uint16 = iota
	// LD 2  ~> load
	LD uint16 = iota
	// ST 3  ~> store
	ST uint16 = iota
	// JSR 4  ~> jump register
	JSR uint16 = iota
	// AND 5  ~> bitwise and
	AND uint16 = iota
	// LDR 6  ~> load register
	LDR uint16 = iota
	// STR 7  ~> store register
	STR uint16 = iota
	// RTI 8  ~> unused
	RTI uint16 = iota
	// NOT 9  ~> bitwise not
	NOT uint16 = iota
	// LDI 10 ~> load indirect
	LDI uint16 = iota
	// STI 11 ~> store indirect
	STI uint16 = iota
	// JMP 12 ~> jump
	JMP uint16 = iota
	// RES 13 ~> reserved (unused)
	RES uint16 = iota
	// LEA 14 ~> load effective address
	LEA uint16 = iota
	// TRAP 15 ~> execute trap
	TRAP uint16 = iota
)

func (cpu *CPU) executeInstruction() {
	// fetch
	var instr uint16 = cpu.memoryRead(cpu.PC)
	var op uint16 = instr >> 12
	cpu.PC++

	switch op {
	case BR:
		cpu.branch(instr)
	case ADD:
		cpu.add(instr)
	case LD:
		cpu.load(instr)
	case ST:
		cpu.store(instr)
	case JSR:
		cpu.jumpRegister(instr)
	case AND:
		cpu.bitwiseAnd(instr)
	case LDR:
		cpu.loadRegister(instr)
	case STR:
		cpu.storeRegister(instr)
	case RTI: // default (unused)
	case NOT:
		cpu.bitwiseNot(instr)
	case LDI:
		cpu.loadIndirect(instr)
	case STI:
		cpu.storeIndirect(instr)
	case JMP:
		cpu.jump(instr)
	case RES: // default (unused)
	case LEA:
		cpu.loadEffectiveAdress(instr)
	case TRAP:
		cpu.trap(instr)
	default: // default (not implemented)
		log.Printf("Operation code not implemented: 0x%04X", instr)
	}
}

//========================================

func (cpu *CPU) branch(instr uint16) {
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	var condFlag = (instr >> 9) & 0x7
	if (condFlag & cpu.COND) != 0 {
		cpu.PC += pcOffset
	}
}

func (cpu *CPU) add(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var immFlag uint16 = (instr >> 5) & 0x1
	if immFlag != 0 {
		var imm5 uint16 = signExtend(instr&0x1F, 5)
		cpu.Registers[r0] = cpu.Registers[r1] + imm5
	} else {
		var r2 uint16 = instr & 0x7
		cpu.Registers[r0] = cpu.Registers[r1] + cpu.Registers[r2]
	}
	cpu.updateFlags(r0)
}

func (cpu *CPU) load(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	cpu.Registers[r0] = cpu.memoryRead(cpu.PC + pcOffset)
	cpu.updateFlags(r0)
}

func (cpu *CPU) store(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	cpu.memoryWrite(cpu.PC+pcOffset, cpu.Registers[r0])
}

func (cpu *CPU) jumpRegister(instr uint16) {
	var longFlag uint16 = (instr >> 11) & 1
	cpu.Registers[7] = cpu.PC
	if longFlag != 0 {
		var longPcOffset uint16 = signExtend(instr&0x7FF, 11)
		cpu.PC += longPcOffset
	} else {
		var r1 uint16 = (instr >> 6) & 0x7
		cpu.PC = cpu.Registers[r1]
	}
}

func (cpu *CPU) bitwiseAnd(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var immFlag uint16 = (instr >> 5) & 0x1
	if immFlag != 0 {
		var imm5 uint16 = signExtend(instr&0x1F, 5)
		cpu.Registers[r0] = cpu.Registers[r1] & imm5
	} else {
		var r2 uint16 = instr & 0x7
		cpu.Registers[r0] = cpu.Registers[r1] & cpu.Registers[r2]
	}
}

func (cpu *CPU) loadRegister(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var offset uint16 = signExtend(instr&0x3F, 6)
	cpu.Registers[r0] = cpu.memoryRead(cpu.Registers[r1] + offset)
	cpu.updateFlags(r0)
}

func (cpu *CPU) storeRegister(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	var offset uint16 = signExtend(instr&0x3F, 6)
	cpu.memoryWrite(cpu.Registers[r1]+offset, cpu.Registers[r0])
}

func (cpu *CPU) bitwiseNot(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var r1 uint16 = (instr >> 6) & 0x7
	cpu.Registers[r0] = ^cpu.Registers[r1]
	cpu.updateFlags(r0)
}

func (cpu *CPU) loadIndirect(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	cpu.Registers[r0] = cpu.memoryRead(cpu.memoryRead(cpu.PC + pcOffset))
	cpu.updateFlags(r0)
}

func (cpu *CPU) storeIndirect(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	cpu.memoryWrite(cpu.memoryRead(cpu.PC+pcOffset), cpu.Registers[r0])
}

func (cpu *CPU) jump(instr uint16) {
	var r1 uint16 = (instr >> 6) & 0x7
	cpu.PC = cpu.Registers[r1]
}

func (cpu *CPU) loadEffectiveAdress(instr uint16) {
	var r0 uint16 = (instr >> 9) & 0x7
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	cpu.Registers[r0] = cpu.PC + pcOffset
	cpu.updateFlags(r0)
}

//========================================

func signExtend(x uint16, bitCount int) uint16 {
	if ((x >> (bitCount - 1)) & 1) != 0 {
		x |= (0xFFFF << bitCount)
	}
	return x
}

func (cpu *CPU) updateFlags(regIndex uint16) {
	if cpu.Registers[regIndex] == 0 {
		cpu.COND = FLGZRO
	} else if cpu.Registers[regIndex]>>15 != 0 { // a 1 in the left-most bit indicates negative
		cpu.COND = FLGNEG
	} else {
		cpu.COND = FLGPOS
	}
}

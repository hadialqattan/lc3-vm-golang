package vm

// Flags
const (
	// FLGPOS for 'Positive' result
	FLGPOS uint16 = 1 << 0
	// FLGZRO for 'Zero' result
	FLGZRO uint16 = 1 << 1
	// FLGNEG for 'Negative' result
	FLGNEG uint16 = 1 << 2
)

type opcodes struct {
	BR   uint16 // 0  ~> branch
	ADD  uint16 // 1  ~> add
	LD   uint16 // 2  ~> load
	ST   uint16 // 3  ~> store
	JSR  uint16 // 4  ~> jump register
	AND  uint16 // 5  ~> bitwise and
	LDR  uint16 // 6  ~> load register
	STR  uint16 // 7  ~> store register
	RTI  uint16 // 8  ~> unused
	NOT  uint16 // 9  ~> bitwise not
	LDI  uint16 // 10 ~> load indirect
	STI  uint16 // 11 ~> store indirect
	JMP  uint16 // 12 ~> jump
	RES  uint16 // 13 ~> reserved (unused)
	LEA  uint16 // 14 ~> load effective address
	TRAP uint16 // 15 ~> execute trap
}

func (cpu *CPU) executeInstruction() {
	// FETCH
	var instr uint16 = cpu.memoryRead(cpu.Registers[cpu.PC])
	var op uint16 = instr >> 12
	cpu.Registers[cpu.PC]++

	switch op {
	case cpu.Opcodes.BR:
		cpu.branch(instr)
	case cpu.Opcodes.ADD:
		cpu.add(instr)
	case cpu.Opcodes.LD:
		cpu.load(instr)
	case cpu.Opcodes.ST:
		cpu.store(instr)
	case cpu.Opcodes.JSR:
		cpu.jumpRegister(instr)
	case cpu.Opcodes.AND:
		cpu.bitwiseAnd(instr)
	case cpu.Opcodes.LDR:
		cpu.loadRegister(instr)
	case cpu.Opcodes.STR:
		cpu.storeRegister(instr)
	case cpu.Opcodes.RTI: // default (unused)
	case cpu.Opcodes.NOT:
		cpu.bitwiseNot(instr)
	case cpu.Opcodes.LDI:
		cpu.loadIndirect(instr)
	case cpu.Opcodes.STI:
		cpu.storeIndirect(instr)
	case cpu.Opcodes.JMP:
		cpu.jump(instr)
	case cpu.Opcodes.RES: // default (unused)
	case cpu.Opcodes.LEA:
		cpu.loadEffectiveAdress(instr)
	case cpu.Opcodes.TRAP:
		cpu.trap(instr)
	default: // default (unused)
	}
}

//========================================

func (cpu *CPU) branch(instr uint16) {
	var pcOffset uint16 = signExtend(instr&0x1FF, 9)
	var condFlag = (instr >> 9) & 0x7
	if condFlag&cpu.Registers[cpu.COND] != 0 {
		cpu.Registers[cpu.PC] += pcOffset
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
		cpu.Registers[cpu.PC] += longPcOffset
	} else {
		var r1 uint16 = (instr >> 6) & 0x7
		cpu.Registers[cpu.PC] = cpu.Registers[r1]
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
	cpu.Registers[r0] = cpu.memoryRead(cpu.PC + pcOffset)
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
	if (x>>(bitCount-1))&1 != 0 {
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

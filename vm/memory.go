package vm

import (
	"encoding/binary"
	"io/ioutil"
	"log"

	termbox "github.com/nsf/termbox-go"
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

//====================================================

// LoadProgramImage from the given path
func (cpu *CPU) LoadProgramImage(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Can't load %s: %s", path, err)
		return err
	}

	// load into the CPU memory
	origin := binary.BigEndian.Uint16(bytes[:2])
	for i := 2; i < len(bytes); i += 2 {
		cpu.Memory[origin] = binary.BigEndian.Uint16(bytes[i : i+2])
		origin++
	}
	return nil
}

// KBinputLoop is a keyboard input processing loop
func (cpu *CPU) KBinputLoop() {
	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			cpu.KeysBuffer = append(cpu.KeysBuffer, event.Ch)
			switch {
			case event.Ch == 'q' || event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC || event.Key == termbox.KeyCtrlD:
				// terminate the CPU processes
				cpu.Stop()
				return
			}
		}
	}
}

func (cpu *CPU) processKBInput() {
	valOfKBSR := cpu.memoryRead(KBSR)
	isKBSRready := ((valOfKBSR & 0x8000) == 0)
	if isKBSRready && len(cpu.KeysBuffer) > 0 {
		cpu.memoryWrite(KBSR, valOfKBSR|0x8000)
		cpu.memoryWrite(KBDR, uint16(cpu.KeysBuffer[0]))
	}
}

package main

import (
	"flag"
	"log"
	"os"

	"github.com/hadialqattan/lc3-vm-golang/vm"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	// init text-based user interface
	err := termbox.Init()
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	defer termbox.Close()

	// get the program file path
	progPath := getProgramPath()

	// init the CPU
	log.Println("Booting the LC3-VM...")
	termbox.Flush()
	cpu := vm.NewCPU()

	// load the program file
	log.Printf("Loading the program: %s", progPath)
	cpu.LoadProgramImage(progPath)

	// start the input loop as goroutine
	go cpu.ProcessKBInput()

	// start processing
	cpu.Run()
	log.Println("Terminated!")

}

func getProgramPath() string {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Printf("A program file must be specified")
		os.Exit(1)
	}
	if info, err := os.Stat(args[0]); err != nil {
		log.Printf("No program file found")
		os.Exit(1)
	} else if info.IsDir() {
		log.Printf("A program should be a file not a directory")
		os.Exit(1)
	}
	return args[0]
}

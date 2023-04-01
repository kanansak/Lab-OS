package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	name   string
	size   int
	offset int
}

const (
	MEM_SIZE = 1000
)

var (
	memory [MEM_SIZE]int
	procs  []*Process
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		showMemoryMap()
		fmt.Print("Command: ")
		scanner.Scan()
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "new":
			if len(fields) < 3 {
				fmt.Println("Invalid command")
				continue
			}
			name := fields[1]
			size, err := strconv.Atoi(fields[2])
			if err != nil {
				fmt.Println("Invalid size")
				continue
			}
			if size > MEM_SIZE {
				fmt.Println("Not enough memory")
				continue
			}
			if allocateMemory(name, size) {
				fmt.Println("Process created")
			} else {
				fmt.Println("Not enough memory")
			}
		case "terminate":
			if len(fields) < 2 {
				fmt.Println("Invalid command")
				continue
			}
			name := fields[1]
			if deallocateMemory(name) {
				fmt.Println("Process terminated")
			} else {
				fmt.Println("Process not found")
			}
		/*case "show":
		showMemoryMap()*/
		case "exit":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

func allocateMemory(name string, size int) bool {
	for i := 0; i < MEM_SIZE; i++ {
		if memory[i] == 0 {
			j := i + 1
			for ; j < i+size && j < MEM_SIZE; j++ {
				if memory[j] != 0 {
					break
				}
			}
			if j == i+size {
				proc := &Process{name, size, i}
				procs = append(procs, proc)
				for k := i; k < j; k++ {
					memory[k] = 1
				}
				return true
			}
		}
	}
	return false
}

func deallocateMemory(name string) bool {
	for i, proc := range procs {
		if proc.name == name {
			for j := proc.offset; j < proc.offset+proc.size; j++ {
				memory[j] = 0
			}
			procs = append(procs[:i], procs[i+1:]...)
			return true
		}
	}
	return false
}

func showMemoryMap() {
	fmt.Println("Name\tStart\tSize")
	for _, proc := range procs {
		fmt.Printf("%s\t%d\t%d\n", proc.name, proc.offset, proc.size)
	}
}

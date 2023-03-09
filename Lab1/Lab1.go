package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1  string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func initialized() {
	cpu1 = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU1   -> %s\n", cpu1)
	fmt.Printf("CPU2   -> %s\n", cpu2)
	fmt.Printf("Ready -> ")
	for i := range ready {
		if ready[i] != "" {
			fmt.Printf("%s ", ready[i])
		}
	}
	fmt.Printf("\n")
	if len(io1) > 0 {
		fmt.Printf("I/O 1 -> ")
		for i := range io1 {
			if io1[i] != "" {
				fmt.Printf("%s ", io1[i])
			}
		}
	}
	fmt.Printf("\n")
	if len(io2) > 0 {
		fmt.Printf("I/O 2 -> ")
		for i := range io2 {
			if io2[i] != "" {
				fmt.Printf("%s ", io2[i])
			}
		}
	}
	fmt.Printf("\n")
	if len(io3) > 0 {
		fmt.Printf("I/O 3 -> ")
		for i := range io3 {
			if io3[i] != "" {
				fmt.Printf("%s ", io3[i])
			}
		}
	}
	fmt.Printf("\n")
	if len(io4) > 0 {
		fmt.Printf("I/O 4 -> ")
		for i := range io4 {
			if io4[i] != "" {
				fmt.Printf("%s ", io4[i])
			}
		}
	}
	fmt.Printf("\n\nCommand > ")

}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_new(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}
func command_terminate() {
	if cpu1 != "" {
		cpu1 = deleteQueue(ready)
	} else if cpu2 != "" {
		cpu2 = deleteQueue(ready)
	}
}

func command_expire() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io1() {
	if cpu1 != "" {
		insertQueue(io1, cpu1)
		cpu1 = ""
	} else if cpu2 != "" {
		insertQueue(io1, cpu2)
		cpu2 = ""
	}
	command_expire()
}

func command_io2() {
	if cpu1 != "" {
		insertQueue(io2, cpu1)
		cpu1 = ""
	} else if cpu2 != "" {
		insertQueue(io2, cpu2)
		cpu2 = ""
	}
	command_expire()
}
func command_io3() {
	if cpu1 != "" {
		insertQueue(io3, cpu1)
		cpu1 = ""
	} else if cpu2 != "" {
		insertQueue(io3, cpu2)
		cpu2 = ""
	}
	command_expire()
}

func command_io4() {
	if cpu1 != "" {
		insertQueue(io4, cpu1)
		cpu1 = ""
	} else if cpu2 != "" {
		insertQueue(io4, cpu2)
		cpu2 = ""
	}
	command_expire()
}
func command_io1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}
func command_io3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}
func command_io4x() {
	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new(commandx[i])
			}
		case "terminate":
			command_terminate()
		case "expire":
			command_expire()
		case "io1":
			command_io1()
		case "io2":
			command_io2()
		case "io3":
			command_io3()
		case "io4":
			command_io4()
		case "io1x":
			command_io1x()
		case "io2x":
			command_io2x()
		case "io3x":
			command_io3x()
		case "io4x":
			command_io4x()
		default:
			fmt.Printf("\nSorry !!! Command Error !!!\n")
		}
	}
}

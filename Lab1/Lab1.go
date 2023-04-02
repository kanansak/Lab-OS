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

// เริ่มต้นตัวแปรโกลบอลเป็นค่าเริ่มต้น
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
	fmt.Printf("\n------------------\n")
	fmt.Printf("CPU_1 -> %s\n", cpu1)
	fmt.Printf("CPU_2 -> %s\n", cpu2)
	fmt.Printf("Ready -> ")
	for i := range ready {
		fmt.Printf("%s ", ready[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O_1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O_2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O_3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O_4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n\nCommand > ")
}

// ใช้สำหรับรับค่าข้อมูลที่ผู้ใช้กรอกเข้ามาผ่านทาง command line interface (CLI)
func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

// new <process_name>: creates a new process with the given name and adds it to the system.
func command_new(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

// terminate <cpu_name>: terminates the process running on the specified CPU core (cpu1 or cpu2).
func command_terminate(namecpu string) {
	if namecpu == "cpu1" {
		if cpu1 != "" {
			cpu1 = deleteQueue(ready)
		}
	}
	if namecpu == "cpu2" {
		if cpu2 != "" {
			cpu2 = deleteQueue(ready)
		}
	}
}

/*func command_terminate_cpu2() {
	if cpu2 != "" {
		cpu2 = deleteQueue(ready)
	}
}*/
//expire cpu1 or expire cpu2: moves the process running on the specified CPU core to the end of the ready queue.
func command_expire_cpu1() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	insertQueue(ready, cpu1)
	cpu1 = p
}
func command_expire_cpu2() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	insertQueue(ready, cpu2)
	cpu2 = p
}

// io1 cpu1 or io1 cpu2: moves the process running on the specified CPU core to the I/O queue 1.
func command_io1_cpu1() {
	insertQueue(io1, cpu1)
	cpu1 = ""
	command_expire_cpu1()
}
func command_io1_cpu2() {
	insertQueue(io1, cpu2)
	cpu2 = ""
	command_expire_cpu2()
}

// io2 cpu1 or io2 cpu2: moves the process running on the specified CPU core to the I/O queue 2.
func command_io2_cpu1() {
	insertQueue(io2, cpu1)
	cpu1 = ""
	command_expire_cpu1()
}
func command_io2_cpu2() {
	insertQueue(io2, cpu2)
	cpu2 = ""
	command_expire_cpu2()
}

// io3 cpu1 or io3 cpu2: moves the process running on the specified CPU core to the I/O queue 3.
func command_io3_cpu1() {
	insertQueue(io3, cpu1)
	cpu1 = ""
	command_expire_cpu1()
}
func command_io3_cpu2() {
	insertQueue(io3, cpu2)
	cpu2 = ""
	command_expire_cpu2()
}

// io4 cpu1 or io4 cpu2: moves the process running on the specified CPU core to the I/O queue 4.
func command_io4_cpu1() {
	insertQueue(io4, cpu1)
	cpu1 = ""
	command_expire_cpu1()
}
func command_io4_cpu2() {
	insertQueue(io4, cpu2)
	cpu2 = ""
	command_expire_cpu2()
}

// io1x, io2x, io3x, io4x: moves the first process from the specified I/O queue to the ready queue if there is an available CPU core.
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

// เป็นฟังก์ชันที่ใช้สำหรับการเพิ่มข้อมูลเข้าไปในคิว (Queue)
func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

// ใช้สำหรับลบข้อมูลจากตัวแปร
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
			command_terminate(commandx[1])
		//case "terminate2":
		//	command_terminate_cpu2()
		case "expire1":
			command_expire_cpu1()
		case "expire2":
			command_expire_cpu2()
		case "io1cpu1":
			command_io1_cpu1()
		case "io1cpu2":
			command_io1_cpu2()
		case "io2cpu1":
			command_io2_cpu1()
		case "io2cpu2":
			command_io2_cpu2()
		case "io3cpu1":
			command_io3_cpu1()
		case "io3cpu2":
			command_io3_cpu2()
		case "io4cpu1":
			command_io4_cpu1()
		case "io4cpu2":
			command_io4_cpu2()
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

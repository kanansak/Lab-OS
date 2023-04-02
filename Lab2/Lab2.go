package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1 string
	cpu2 string

	ready1 []string
	ready2 []string
	ready3 []string

	io1 []string
	io2 []string
	io3 []string
	io4 []string

	pty_cpu1 string
	//pty_cpu1 string

	pty_ready1 []string
	pty_ready2 []string
	pty_ready3 []string

	pty_io1 []string
	pty_io2 []string
	pty_io3 []string
	pty_io4 []string

	q1 int
	q2 int
	q3 int
)

func initialized() {
	cpu1 = ""
	cpu2 = ""

	ready1 = make([]string, 10)
	ready2 = make([]string, 10)
	ready3 = make([]string, 10)

	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)

	pty_cpu1 = ""
	pty_cpu1 = ""

	pty_ready1 = make([]string, 10)
	pty_ready2 = make([]string, 10)
	pty_ready3 = make([]string, 10)

	pty_io1 = make([]string, 10)
	pty_io2 = make([]string, 10)
	pty_io3 = make([]string, 10)
	pty_io4 = make([]string, 10)

	q1 = 0
	q2 = 0
	q3 = 0
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU_1 -> %s \n", cpu1)
	fmt.Printf("CPU_2 -> %s \n", cpu2)
	fmt.Printf("Ready_1 [%d]-> ", q1)
	for i := range ready1 {
		fmt.Printf("%s ", ready1[i])
	}
	fmt.Printf("\nReady_2 [%d]-> ", q2)
	for i := range ready2 {
		fmt.Printf("%s ", ready2[i])
	}
	fmt.Printf("\nReady_3 -> ")
	for i := range ready3 {
		fmt.Printf("%s ", ready3[i])
	}
	fmt.Printf("\nI/O_1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\nI/O_2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\nI/O_3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\nI/O_4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\nCommand: ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func insertQueue(q []string, data string, qplt []string, cout_p string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			qplt[i] = cout_p
			break
		}
	}
}

func command_new(p string, cout_p string) {
	if cpu1 == "" {
		cpu1 = p
		pty_cpu1 = cout_p
		if cout_p == "1" {
			q1++
		} else if cout_p == "2" {
			q2++
		} else if cout_p == "3" {
			q3++
		}
	} else if cpu2 == "" {
		cpu2 = p
		pty_cpu1 = cout_p
		if cout_p == "1" {
			q1++
		} else if cout_p == "2" {
			q2++
		} else if cout_p == "3" {
			q3++
		}

	} else {
		if cout_p == "1" {
			insertQueue(ready1, p, pty_ready1, cout_p)
		} else if cout_p == "2" {
			insertQueue(ready2, p, pty_ready2, cout_p)
		} else if cout_p == "3" {
			insertQueue(ready3, p, pty_ready3, cout_p)
		}
	}
}
func deleteQueue(q []string, cout_p []string) (string, string) {
	result := q[0]
	resultp := cout_p[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
		cout_p[i-1] = cout_p[i]
	}
	q[9] = ""
	cout_p[9] = ""
	return result, resultp
}
func command_expire(cpuName string) {
	if cpuName == "cpu1" {
		cout_p := pty_cpu1
		CheckExpireCpu1(cout_p)
	} else if cpuName == "cpu2" {
		cout_p := pty_cpu1
		CheckExpireCpu2(cout_p)
	}
	newQueue := ""
	newPiority := ""
	if q1 < 3 && ready1[0] != "" {
		newQueue, newPiority = deleteQueue(ready1, pty_ready1)
	} else if q2 < 3 && ready2[0] != "" {
		newQueue, newPiority = deleteQueue(ready2, pty_ready2)
		if q2 < 2 {
			if q1 == 3 {
				q1 = 0
				if q2 == 3 {
					q2 = 0
				}
			} else if q2 == 3 {
				q2 = 0
			} else if q3 == 3 {
				q3 = 0
			}
		}
	} else if q3 < 3 && ready3[0] != "" {
		newQueue, newPiority = deleteQueue(ready3, pty_ready3)
		if q1 == 3 {
			q1 = 0
			if q2 == 3 {
				q2 = 0
			}
		} else if q2 == 3 {
			q2 = 0
		} else if q3 == 3 {
			q3 = 0
		}
	}
	if newPiority == "1" {
		q1++
	} else if newPiority == "2" {
		q2++
	} else if newPiority == "3" {
		q3++
	}
	if newQueue == "" {
		return
	}

	if cpuName == "cpu1" {
		cpu1 = newQueue
		pty_cpu1 = newPiority
	} else if cpuName == "cpu2" {
		cpu2 = newQueue
		pty_cpu1 = newPiority
	}
}

func CheckExpireCpu1(inputPiorCpu1 string) {
	if inputPiorCpu1 == "1" {
		insertQueue(ready1, cpu1, pty_ready1, pty_cpu1)
	} else if inputPiorCpu1 == "2" {
		insertQueue(ready2, cpu1, pty_ready2, pty_cpu1)
	} else if inputPiorCpu1 == "3" {
		insertQueue(ready3, cpu1, pty_ready3, pty_cpu1)
	}
}

func CheckExpireCpu2(inputPiorCpu2 string) {
	if inputPiorCpu2 == "1" {
		insertQueue(ready1, cpu2, pty_ready1, pty_cpu1)
	} else if inputPiorCpu2 == "2" {
		insertQueue(ready2, cpu2, pty_ready2, pty_cpu1)
	} else if inputPiorCpu2 == "3" {
		insertQueue(ready3, cpu2, pty_ready3, pty_cpu1)
	}
}
func command_ioS(ioName string, cpuName string) {
	switch ioName {
	case "1":
		io_cpu(io1, pty_io1, cpuName)
	case "2":
		io_cpu(io2, pty_io2, cpuName)
	case "3":
		io_cpu(io3, pty_io3, cpuName)
	case "4":
		io_cpu(io4, pty_io4, cpuName)
	default:
		return
	}
}

func io_cpu(io []string, iop []string, cpu string) {
	if cpu == "cpu1" {
		insertQueue(io, cpu1, iop, pty_cpu1)
		cpu1 = ""
		pty_cpu1 = ""
	} else if cpu == "cpu2" {
		insertQueue(io, cpu2, iop, pty_cpu1)
		cpu2 = ""
		pty_cpu1 = ""
	}
	command_expire(cpu)
}
func command_terminate(cpuName string) {
	if cpuName == "cpu1" {
		if q1 < 3 && ready1[0] != "" {
			cpu1, pty_cpu1 = deleteQueue(ready1, pty_ready1)
		} else if q2 < 3 && ready2[0] != "" {
			cpu1, pty_cpu1 = deleteQueue(ready2, pty_ready2)
		} else if q3 < 3 && ready3[0] != "" {
			cpu1, pty_cpu1 = deleteQueue(ready3, pty_ready3)
		} else if ready1[0] == "" && ready2[0] == "" && ready3[0] == "" {
			cpu1 = ""
			pty_cpu1 = ""
		}
		if q1 == 3 {
			q1 = 0
			if q2 == 3 {
				q2 = 0
			}
		} else if q2 == 3 {
			q2 = 0
		} else if q3 == 3 {
			q3 = 0
		}

		if pty_cpu1 == "1" {
			q1++
		} else if pty_cpu1 == "2" {
			q2++
		} else if pty_cpu1 == "3" {
			q3++
		}
	} else if cpuName == "cpu2" {
		if q1 < 3 && ready1[0] != "" {
			cpu2, pty_cpu1 = deleteQueue(ready1, pty_ready1)
		} else if q2 < 3 && ready2[0] != "" {
			cpu2, pty_cpu1 = deleteQueue(ready2, pty_ready2)
		} else if q3 < 3 && ready3[0] != "" {
			cpu2, pty_cpu1 = deleteQueue(ready3, pty_ready3)
		} else if ready1[0] == "" && ready2[0] == "" && ready3[0] == "" {
			cpu2 = ""
			pty_cpu1 = ""
		}
		if q1 == 3 {
			q1 = 0
			if q2 == 3 {
				q2 = 0
			}
		} else if q2 == 3 {
			q2 = 0
		} else if q3 == 3 {
			q3 = 0
		}

		if pty_cpu1 == "1" {
			q1++
		} else if pty_cpu1 == "2" {
			q2++
		} else if pty_cpu1 == "3" {
			q3++
		}
	}
}
func command_ioSx(ioName string) {
	fq := ""
	cout_p := ""
	switch ioName {
	case "1":
		fq, cout_p = deleteQueue(io1, pty_io1)
	case "2":
		fq, cout_p = deleteQueue(io2, pty_io2)
	case "3":
		fq, cout_p = deleteQueue(io3, pty_io3)
	case "4":
		fq, cout_p = deleteQueue(io4, pty_io4)
	default:
		return
	}
	if fq == "" {
		return
	}

	if cpu1 == "" {
		cpu1 = fq
		pty_cpu1 = cout_p

		if cout_p == "1" {
			q1++
		} else if cout_p == "2" {
			q2++
		} else if cout_p == "3" {
			q3++
		}
	} else if cpu2 == "" {
		cpu2 = fq
		pty_cpu1 = cout_p

		if cout_p == "1" {
			q1++
		} else if cout_p == "2" {
			q2++
		} else if cout_p == "3" {
			q3++
		}
	} else {
		if cout_p == "1" {
			insertQueue(ready1, fq, pty_ready1, cout_p)
		} else if cout_p == "2" {
			insertQueue(ready2, fq, pty_ready2, cout_p)
		} else if cout_p == "3" {
			insertQueue(ready3, fq, pty_ready3, cout_p)
		}
	}
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
				if i%2 != 0 {
					command_new(commandx[i], commandx[i+1])
				}

			}
		case "ter":
			command_terminate(commandx[1])
		case "exp":
			command_expire(commandx[1])
		case "io":
			command_ioS(commandx[1], commandx[2])
		case "iox":
			command_ioSx(commandx[1])
		default:
			fmt.Printf("\nInput Error \n")
		}
	}

}

//new p1 1

//io 1 cpu1

//iox 1

//exp cpu1
//exp cpu2

//ter cpu1
//ter cpu2

//new p1 1 p2 1 p3 1 p4 2 p5 3
//exp cpu1
//exp cpu2

//exp cpu1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	process   []string
	allocate  []int
	need      []int
	max       []int
	available []int
)

func initialized() {
	process = make([]string, 10)
	allocate = make([]int, 30)
	need = make([]int, 30)
	max = make([]int, 30)
	available = make([]int, 3)

	for i := range available {
		available[i] = 10
	}
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func showProcess() {
	fmt.Printf("\n+-----------------------------------------------+\n")
	fmt.Printf(" Process |Allocate|  Need |  Max  | Available ")
	fmt.Printf("\n         | A B C  | A B C | A B C | ")
	fmt.Printf("\n+-----------------------------------------------+\n")

	if process[0] == "" {
		fmt.Printf("    -    | - - -  | - - - | - - - | %d %d %d\n", available[0], available[1], available[2])
	} else {
		for i := range process {
			if process[i] == "" {
				continue
			} else {
				if i == 0 {
					fmt.Printf("    %s   | %d %d %d  | %d %d %d | %d %d %d | %d %d %d\n", process[i], allocate[0], allocate[1], allocate[2], need[0], need[1], need[2], max[0], max[1], max[2], available[0], available[1], available[2])
				} else {
					fmt.Printf("    %s   | %d %d %d  | %d %d %d | %d %d %d |\n", process[i], allocate[0+(3*i)], allocate[1+(3*i)], allocate[2+(3*i)], need[0+(3*i)], need[1+(3*i)], need[2+(3*i)], max[0+(3*i)], max[1+(3*i)], max[2+(3*i)])
				}
			}
		}
	}
	fmt.Printf("\n")
	fmt.Printf("\nCommand > ")
}

func command_new(p string, m1, m2, m3 int) {
	for i := range process {
		if process[i] == "" {
			process[i] = p
			max[0+(i*3)] = m1
			max[1+(i*3)] = m2
			max[2+(i*3)] = m3
			for i := range process {
				if process[i] == "" {
					continue
				} else {
					need[0+(i*3)] = max[0+(i*3)] - allocate[0+(i*3)]
					need[1+(i*3)] = max[1+(i*3)] - allocate[1+(i*3)]
					need[2+(i*3)] = max[2+(i*3)] - allocate[2+(i*3)]
					if (need[0+(i*3)] == 0) && (need[1+(i*3)] == 0) && (need[2+(i*3)] == 0) {
						command_terminate(i)
					}
				}
			}
			break
		}
	}
}

func command_update() {
	for i := range process {
		if process[i] == "" {
			continue
		} else {
			need[0+(i*3)] = max[0+(i*3)] - allocate[0+(i*3)]
			need[1+(i*3)] = max[1+(i*3)] - allocate[1+(i*3)]
			need[2+(i*3)] = max[2+(i*3)] - allocate[2+(i*3)]
			if (need[0+(i*3)] == 0) && (need[1+(i*3)] == 0) && (need[2+(i*3)] == 0) {
				command_terminate(i)
			}
		}
	}
}

func command_request(p string, a, b, c int) {
	if (available[0]-a > 0) && (available[1]-b > 0) && (available[2]-c > 0) {
		test1 := available[0] - a
		test2 := available[1] - b
		test3 := available[2] - c
		safe := false

		for i := range process {
			if process[i] == "" {
				continue
			} else if process[i] != p {
				if (test1 >= need[0+(i*3)]) && (test2 >= need[1+(i*3)]) && (test3 >= need[2+(i*3)]) {
					safe = true
					break
				}
			} else {
				if (test1 >= (need[0+(i*3)] - a)) && (test2 >= (need[1+(i*3)] - b)) && (test3 >= (need[2+(i*3)] - c)) {
					safe = true
					break
				}
			}
		}

		for i := range process {
			if process[i] == p {
				if (a <= need[0+(i*3)]) && (b <= need[1+(i*3)]) && (c <= need[2+(i*3)]) && safe == true {
					allocate[0+(i*3)] += a
					allocate[1+(i*3)] += b
					allocate[2+(i*3)] += c
					available[0] -= a
					available[1] -= b
					available[2] -= c
					fmt.Printf("\n--------------------Safe!--------------------\n")
					safe = false
				} else {
					fmt.Printf("\n--------------------Not Safe!--------------------\n")
				}
			} else {
				continue
			}
		}
		command_update()
	} else if (available[0]-a == 0) && (available[1]-b == 0) && (available[2]-c == 0) {
		test1 := available[0] - a
		test2 := available[1] - b
		test3 := available[2] - c
		safe := false

		for i := range process {
			if process[i] == "" {
				continue
			} else if process[i] != p {
				if (test1 >= need[0+(i*3)]) && (test2 >= need[1+(i*3)]) && (test3 >= need[2+(i*3)]) {
					safe = true
					break
				}
			} else {
				if (test1 >= (need[0+(i*3)] - a)) && (test2 >= (need[1+(i*3)] - b)) && (test3 >= (need[2+(i*3)] - c)) {
					safe = true
					break
				}
			}
		}

		for i := range process {
			if process[i] == p {
				if (available[0]-need[0+(i*3)] == 0) && (available[1]-need[1+(i*3)] == 0) && (available[2]-need[2+(i*3)] == 0) && safe == true {
					allocate[0+(i*3)] += a
					allocate[1+(i*3)] += b
					allocate[2+(i*3)] += c
					available[0] -= a
					available[1] -= b
					available[2] -= c
					fmt.Printf("\n--------------------Safe!--------------------\n")
					safe = false
				} else {
					fmt.Printf("\n--------------------Not Safe!--------------------\n")
				}
			} else {
				continue
			}
		}
		command_update()
	} else {
		fmt.Printf("\n--------------------Not Safe!--------------------\n")
	}
}

func command_terminate(p int) {
	available[0] += allocate[0+(p*3)]
	available[1] += allocate[1+(p*3)]
	available[2] += allocate[2+(p*3)]
	for i := range process {
		if process[i] == "" {
			break
		}
		if process[i] != process[p] {
			continue
		}
		process[i] = process[i+1]
		need[0+(i*3)] = need[0+(i*3)+3]
		need[1+(i*3)] = need[1+(i*3)+3]
		need[2+(i*3)] = need[2+(i*3)+3]
		max[0+(i*3)] = max[0+(i*3)+3]
		max[1+(i*3)] = max[1+(i*3)+3]
		max[2+(i*3)] = max[2+(i*3)+3]
		allocate[0+(i*3)] = allocate[0+(i*3)+3]
		allocate[1+(i*3)] = allocate[1+(i*3)+3]
		allocate[2+(i*3)] = allocate[2+(i*3)+3]
		p = i + 1
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
			m1, _ := strconv.Atoi(commandx[2])
			m2, _ := strconv.Atoi(commandx[3])
			m3, _ := strconv.Atoi(commandx[4])
			command_new(commandx[1], m1, m2, m3)
		case "req":
			a, _ := strconv.Atoi(commandx[2])
			b, _ := strconv.Atoi(commandx[3])
			c, _ := strconv.Atoi(commandx[4])
			command_request(commandx[1], a, b, c)
		}

	}
}

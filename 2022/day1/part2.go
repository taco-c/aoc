package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var max = [3]int{0, 0, 0}
	var elf = 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if n, err := strconv.Atoi(s.Text()); err == nil {
			elf += n
			continue
		}
		Order(&max, elf)
		elf = 0
	}

	Order(&max, elf)
	var maxInt int
	for _, n := range max {
		maxInt += n
	}
	fmt.Println(maxInt)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Order(max *[3]int, elf int) {
	if elf > max[0] {
		max[2] = max[1]
		max[1] = max[0]
		max[0] = elf
	} else if elf > max[1] {
		max[2] = max[1]
		max[1] = elf
	} else if elf > max[2] {
		max[2] = elf
	}
}

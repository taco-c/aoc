package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	max := 0
	elf := 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if n, err := strconv.Atoi(s.Text()); err == nil {
			elf += n
		} else {
			max = Max(max, elf)
			elf = 0
		}
	}

	max = Max(max, elf)
	fmt.Println(max)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

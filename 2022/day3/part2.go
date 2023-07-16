package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	score := 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		elf1 := s.Text()
		s.Scan()
		elf2 := s.Text()
		s.Scan()
		elf3 := s.Text()

		badge := getCommonItem(elf1, elf2, elf3)

		if badge > 'a' {
			score += int(badge) - 96
		} else {
			score += int(badge) - 38
		}
	}

	fmt.Println(score)
}

func getCommonItem(a, b, c string) rune {
	for _, i1 := range a {
		for _, i2 := range b {
			for _, i3 := range c {
				if i1 == i2 && i1 == i3 {
					return i1
				}
			}
		}
	}
	return 0
}

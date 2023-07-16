package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var score = 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		t := s.Text()
		a := strings.Split(t, " ")

		switch a[1] {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		if a[0] == "A" && a[1] == "X" {
			score += 3
		} else if a[0] == "A" && a[1] == "Y" {
			score += 6
		} else if a[0] == "A" && a[1] == "Z" {
			score += 0
		} else if a[0] == "B" && a[1] == "X" {
			score += 0
		} else if a[0] == "B" && a[1] == "Y" {
			score += 3
		} else if a[0] == "B" && a[1] == "Z" {
			score += 6
		} else if a[0] == "C" && a[1] == "X" {
			score += 6
		} else if a[0] == "C" && a[1] == "Y" {
			score += 0
		} else if a[0] == "C" && a[1] == "Z" {
			score += 3
		} else {
			score += 0
		}
	}

	fmt.Println(score)
}

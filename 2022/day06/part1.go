package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)

	var window [4]rune
	for i := 1; s.Scan(); i++ {
		shiftDown(&window, rune(s.Bytes()[0]))

		if i > 3 && allDifferent(&window) {
			fmt.Println(i)
			break
		}
	}
}

func shiftDown(arr *[4]rune, r rune) {
	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = arr[3]
	arr[3] = r
}

func allDifferent(arr *[4]rune) bool {
	return arr[0] != arr[1] &&
		arr[0] != arr[2] &&
		arr[0] != arr[3] &&
		arr[1] != arr[2] &&
		arr[1] != arr[3] &&
		arr[2] != arr[3]
}

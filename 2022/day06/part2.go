package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 14

func main() {
	var s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)

	var window []rune
	for i := 1; s.Scan(); i++ {
		shiftDown(&window, rune(s.Bytes()[0]), size)

		if i > size && allDifferent(&window) {
			fmt.Println(i)
			break
		}
	}
}

func shiftDown(arr *[]rune, r rune, size int) {
	if len(*arr) < size {
		*arr = append(*arr, r)
		return
	}
	for i := 1; i < len(*arr); i++ {
		(*arr)[i-1] = (*arr)[i]
	}
	(*arr)[len(*arr)-1] = r
}

func allDifferent(arr *[]rune) bool {
	var size = len(*arr)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (*arr)[i] == (*arr)[j] {
				return false
			}
		}
	}
	return true
}

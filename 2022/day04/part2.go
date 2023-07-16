package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var score = 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		var values = strings.FieldsFunc(s.Text(), func(r rune) bool {
			return r == ',' || r == '-'
		})

		var numbers [4]int
		for i, value := range values {
			var n, _ = strconv.Atoi(value)
			numbers[i] = n
		}

		var a = numbers[1] >= numbers[2] && numbers[0] <= numbers[3]
		var b = numbers[0] <= numbers[3] && numbers[1] >= numbers[2]
		if a || b {
			score++
		}
	}

	fmt.Println(score)
}

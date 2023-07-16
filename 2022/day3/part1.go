package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	score := 0

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		half := len(s.Text()) / 2
		commonItemType := getCommonItem(s.Text()[:half], s.Text()[half:])

		if commonItemType > 'a' {
			score += int(commonItemType) - 96
		} else {
			score += int(commonItemType) - 38
		}
	}

	fmt.Println(score)
}

func getCommonItem(a, b string) rune {
	for i := 0; i < len(a); i++ {
		if strings.IndexByte(b, a[i]) >= 0 {
			return rune(a[i])
		}
		// Seems to be slighty faster than this:
		// if strings.ContainsRune(b, rune(a[i])) {
		//     return rune(a[i])
		// }
	}
	return 0
}

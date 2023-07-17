package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	Crates []rune
}

func (s *Stack) Push(r rune) {
	s.Crates = append(s.Crates, r)
}

func (s *Stack) Pop() rune {
	var value = s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[:len(s.Crates)-1]
	return value
}

func (s *Stack) String() string {
	var b strings.Builder
	b.Grow(len(s.Crates) * 4)
	for _, c := range s.Crates {
		fmt.Fprintf(&b, "[%c] ", c)
	}
	return b.String()
}

func main() {
	var s = bufio.NewScanner(os.Stdin)
	var stacks []Stack = parseStacks(s)
	for _, stack := range stacks {
		fmt.Println(stack.String())
	}
	fmt.Println()

	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		var n, from, to int
		fmt.Sscanf(s.Text(), "move %d from %d to %d", &n, &from, &to)
		fmt.Printf("[%d->%d (%d)]\n", from, to, n)
		for i := 0; i < n; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	for _, stack := range stacks {
		fmt.Println(stack.String())
	}

	var sb strings.Builder
	for _, stack := range stacks {
		fmt.Fprintf(&sb, "%c", stack.Pop())
	}
	fmt.Println(sb.String())
}

func parseStacks(s *bufio.Scanner) []Stack {
	var stacks []Stack
	for s.Scan() {
		s.Text()

		var line = s.Text()

		// Read 4 runes at a time.
		for n, m, column := 0, 3, 0; m <= len(line); {
			var symbol = line[n:m]
			fmt.Printf("%d -> symbol '%s'\n", column+1, symbol)
			switch symbol[0] {
			case ' ':
				if symbol[1] == '1' {
					goto end
				}
			case '[':
				for len(stacks) < column+1 {
					stacks = append(stacks, Stack{})
				}
				stacks[column].Push(rune(symbol[1]))
			}

			n = m + 1
			m = n + 3
			column++
		}
	}

end:
	for i := 0; i < len(stacks); i++ {
		for n, m := 0, len(stacks[i].Crates)-1; n < m; n, m = n+1, m-1 {
			stacks[i].Crates[n], stacks[i].Crates[m] = stacks[i].Crates[m], stacks[i].Crates[n]
		}
	}

	return stacks
}

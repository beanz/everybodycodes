package main

import (
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

func potions(in []byte, n int) int {
	if in[len(in)-1] == '\n' {
		in = in[:len(in)-1]
	}
	res := 0
	for i := 0; i < len(in); i += n {
		c := []int{0, 0, 0, 0, 0}
		for _, ch := range in[i : i+n] {
			switch ch {
			case 'x':
				c[4]++
			default:
				c[ch-'A']++
			}
		}
		nx := n - c[4]
		if nx == 0 {
			continue
		}
		extra := nx - 1
		s := (5+extra)*c[3] + (3+extra)*c[2] + (1+extra)*c[1] + (0+extra)*c[0]
		res += s
	}
	return res
}

func main() {
	i1 := Input(1, "")
	fmt.Printf("Part 1: %v\n", potions(i1, 1))
	i2 := Input(2, "")
	fmt.Printf("Part 2: %v\n", potions(i2, 2))
	i3 := Input(3, "")
	fmt.Printf("Part 3: %v\n", potions(i3, 3))
}

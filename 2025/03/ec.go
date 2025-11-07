package main

import (
	"fmt"
	"sort"

	. "github.com/beanz/everybodycodes/lib-go"
)

func main() {
	i1 := Input(1, "")
	num := []int{}
	for i := 0; i < len(i1); {
		j, n := ChompUInt[int](i1, i)
		i = j + 1
		num = append(num, n)
	}
	sort.Ints(num)
	p1 := num[0]
	p := num[0]
	for i := 1; i < len(num); i++ {
		if p == num[i] {
			continue
		}
		p = num[i]
		p1 += p
	}
	i2 := Input(2, "")
	num = num[:0]
	for i := 0; i < len(i2); {
		j, n := ChompUInt[int](i2, i)
		i = j + 1
		num = append(num, n)
	}
	sort.Ints(num)
	p2 := num[0]
	p = num[0]
	c := 1
	for i := 1; i < len(num); i++ {
		if p == num[i] {
			continue
		}
		c++
		p = num[i]
		p2 += p
		if c == 20 {
			break
		}
	}
	i3 := Input(3, "")
	num = num[:0]
	for i := 0; i < len(i3); {
		j, n := ChompUInt[int](i3, i)
		i = j + 1
		num = append(num, n)
	}
	sort.Ints(num)
	p3 := 0
	for i := 0; i < len(num); {
		n := num[i]
		c := 0
		for ; i < len(num); i++ {
			if num[i] == n {
				c++
			} else {
				break
			}
		}
		if p3 < c {
			p3 = c
		}
	}
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Part 3: %d\n", p3)
}

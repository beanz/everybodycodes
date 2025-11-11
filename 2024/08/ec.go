package main

import (
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	var a1 int
	{
		_, n := ChompUInt[int](i1, 0)
		t, i := 0, 1
		for ; t < n; i += 2 {
			t += i
		}
		i -= 2
		d := t - n
		a1 = i * d
	}
	var a2 int
	{
		_, priests := ChompUInt[int](i2, 0)
		acolytes, blocks := 1111, 20240000
		if priests == 3 {
			acolytes, blocks = 5, 50
		}
		t, i, th := 1, 3, 1
		for ; t < blocks; i += 2 {
			th = (th * priests) % acolytes
			t += i * th
		}
		i -= 2
		d := t - blocks
		a2 = i * d
	}
	var a3 int
	{
		_, priests := ChompUInt[int](i3, 0)
		acolytes, blocks := 10, 202400000
		if priests == 2 {
			acolytes, blocks = 5, 160
		}
		t, i, th := 1, 3, 1
		heights := []int{1}
		for ; t < blocks; i += 2 {
			th = (th*priests)%acolytes + acolytes
			t = 0
			heights = append(heights, 0)
			for i := range heights {
				heights[i] += th
				t += heights[i] * 2
			}
			t -= heights[0]
		}
		i -= 2
		remove := 0
		m := i * priests
		for i := range len(heights) - 1 {
			r := m * heights[i] % acolytes
			r = min(heights[i+1], r)
			if i > 0 {
				r *= 2
			}
			remove += r
		}
		t -= remove
		d := t - blocks
		a3 = d
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

package main

import (
	"fmt"
	"math"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1, _ := part1(i1, 10)
	_, a2 := part1(i2, math.MaxInt)
	a3 := part3(i3)
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

func part3(in []byte) int {
	var nums []int
	var total int
	for i := 0; i < len(in); i++ {
		var n int
		i, n = ChompUInt[int](in, i)
		nums = append(nums, n)
		total += n
	}
	avg := total / len(nums)
	a3 := 0
	for _, n := range nums {
		if n < avg {
			a3 += avg - n
		}
	}
	return a3
}

func part1(in []byte, rounds int) (int, int) {
	a := 0
	var nums []int
	for i := 0; i < len(in); i++ {
		var n int
		i, n = ChompUInt[int](in, i)
		nums = append(nums, n)
	}
	i := 0
	for ; i < rounds; i++ {
		var phase1 bool
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				phase1 = true
				nums[j]--
				nums[j+1]++
			}
		}
		if !phase1 {
			break
		}
		//Lg("p1", i, nums)
	}
	for ; i < rounds; i++ {
		var phase2 bool
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] < nums[j+1] {
				phase2 = true
				nums[j]++
				nums[j+1]--
			}
		}
		if !phase2 {
			break
		}
		//Lg("p2", i, nums)
	}
	for i, n := range nums {
		a += (i + 1) * n
	}
	return a, i
}

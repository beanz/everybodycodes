package ec

import "slices"

type Perms struct {
	orig []int
	perm []int
}

func NewPerms(n int) *Perms {
	p := &Perms{orig: make([]int, n), perm: make([]int, n)}
	for i := 0; i < n; i++ {
		p.orig[i] = i
	}
	return p
}

func (p *Perms) Done() bool {
	return p.perm[0] == len(p.perm)-1
}

func (p *Perms) Next() []int {
	for i := len(p.perm) - 1; i >= 0; i-- {
		if i == 0 || p.perm[i] < len(p.perm)-i-1 {
			p.perm[i]++
			break
		}
		p.perm[i] = 0
	}
	return p.Get()
}

func (p *Perms) Get() []int {
	result := append([]int{}, p.orig...)
	for i, v := range p.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func Permutations[T any](sets ...[]T) [][]T {
	var perms [][]T
	used := make([]int, len(sets))
	var aux func(cur []T, used ...int)
	aux = func(curr []T, used ...int) {
		done := true
		for i, set := range sets {
			if used[i] < len(set) {
				done = false
				u := slices.Clone(used)
				u[i]++
				aux(append(slices.Clone(curr), set[0]), u...)
			}
		}
		if done {
			perms = append(perms, curr)
		}
	}
	aux(nil, used...)
	return perms
}

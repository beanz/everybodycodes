package main

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, int, int) {
	p1 := parse(i1)
	a1 := p1.names[func() int {
		for j, name := range p1.names {
			if p1.check(name) {
				return j
			}
		}
		return -1
	}()]

	p2 := parse(i2)
	a2 := func() int {
		a := 0
		for j, name := range p2.names {
			if p2.check(name) {
				a += j + 1
			}
		}
		return a
	}()
	p3 := parse(i3)
	a3 := p3.count(7, 11)
	return string(a1), a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	names [][]byte
	rules map[byte][]byte
}

func (p *parsed) count(minLen, maxLen int) int {
	var names [][]byte
NAME:
	for i := 0; i < len(p.names); i++ {
		if !p.check(p.names[i]) {
			continue
		}
		for j := 0; j < len(p.names); j++ {
			if i == j {
				continue
			}
			if bytes.HasPrefix(p.names[i], p.names[j]) {
				continue NAME
			}
		}
		names = append(names, p.names[i])
	}
	var aux func(prefix []byte) int
	aux = func(prefix []byte) int {
		l := len(prefix)
		res := 0
		if l > maxLen {
			return res
		}
		if l >= minLen {
			res++
		}
		next := p.rules[prefix[l-1]]
		for _, ch := range next {
			prefix = append(prefix, ch)
			res += aux(prefix)
			prefix = prefix[:l]
		}
		return res
	}
	res := 0
	for _, base := range names {
		res += aux(base)
	}
	return res
}

func (p *parsed) check(name []byte) bool {
	ch := name[0]
	for i := 1; i < len(name); i++ {
		next, ok := p.rules[ch]
		if !ok {
			return false
		}
		nch := name[i]
		if !slices.Contains(next, nch) {
			return false
		}
		ch = nch
	}
	return true
}

func parse(in []byte) *parsed {
	p := &parsed{rules: map[byte][]byte{}}
	names, rules, ok := bytes.Cut(in, []byte("\n\n"))
	if !ok {
		panic(string(in))
	}
	p.names = bytes.Split(names, []byte{','})
	for rule := range bytes.SplitSeq(rules, []byte{'\n'}) {
		if len(rule) == 0 {
			continue
		}
		rhs, lhs, ok := bytes.Cut(rule, []byte(" > "))
		if !ok {
			panic(string(rule))
		}
		for to := range bytes.SplitSeq(lhs, []byte{','}) {
			p.rules[rhs[0]] = append(p.rules[rhs[0]], to[0])
		}
	}
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	for _, name := range p.names {
		fmt.Fprintf(&sb, "%s, ", name)
	}
	fmt.Fprintln(&sb)
	for f, t := range p.rules {
		fmt.Fprintf(&sb, "%c %s\n", f, t)
	}
	return sb.String()
}

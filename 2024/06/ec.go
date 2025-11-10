package main

import (
	"bytes"
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, string, string) {
	t := parse(i1)
	a1 := t.search(true)
	t = parse(i2)
	a2 := t.search(false)
	t = parse(i3)
	a3 := t.search(false)
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	tree map[string][]string
}

func (p *parsed) search(full bool) string {
	l := map[int][]string{}
	p.find("RR", []string{}, full, l)
	paths := slices.Collect(maps.Keys(l))
	slices.SortFunc(paths, func(a int, b int) int {
		return cmp.Compare(len(l[a]), len(l[b]))
	})
	return l[paths[0]][0]
}

func (p *parsed) find(from string, path []string, full bool, res map[int][]string) {
	if from == "@" {
		key := ""
		for _, e := range path {
			if full {
				key += e
			} else {
				key += string(e[0])
			}
		}
		key += "@"
		res[len(path)] = append(res[len(path)], key)
		return
	}
	path = append(path, from)
	for _, branch := range p.tree[from] {
		if slices.Contains(path, branch) {
			continue
		}
		l := len(path)
		p.find(branch, path, full, res)
		path = path[:l]
	}
}

func parse(in []byte) *parsed {
	tree := map[string][]string{}
	for line := range bytes.SplitSeq(in, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		k, branches, _ := bytes.Cut(line, []byte{':'})
		key := string(k)
		for branch := range bytes.SplitSeq(branches, []byte{','}) {
			tree[key] = append(tree[key], string(branch))
		}
	}
	return &parsed{tree}
}

func (p *parsed) String() string {
	var sb strings.Builder
	return sb.String()
}

package ec

type IdentifierMap[T comparable, U ECInt] struct {
	toID map[T]U
	next U
}

func NewIdentifierMap[T comparable, U ECInt]() *IdentifierMap[T, U] {
	return &IdentifierMap[T, U]{toID: map[T]U{}}
}

func (i *IdentifierMap[T, U]) Size() U {
	return i.next
}

func (i *IdentifierMap[T, U]) Add(e T) U {
	if v, ok := i.toID[e]; ok {
		return v
	}
	i.toID[e] = i.next
	i.next++
	if i.next == U(0) {
		panic("IdentifierMap overflow")
	}
	return i.toID[e]
}

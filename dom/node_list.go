package dom

type NodeList interface {
	// accessors
	GetLength() int32
	// spec-defined functions
	Item(index int32) Node
}

type NodeListImpl []Node

func (n NodeListImpl) GetLength() int32 {
	return int32(len(n))
}

func (n NodeListImpl) Item(index int32) Node {
	return n[int(index)]
}

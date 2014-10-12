package dom

type INodeList interface {
	// accessors
	GetLength() int32
	// spec-defined functions
	Item(index int32) INode
}

type NodeList []INode

func (n NodeList) GetLength() int32 {
	return int32(len(n))
}

func (n NodeList) Item(index int32) INode {
	return n[int(index)]
}

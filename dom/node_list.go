package dom

type INodeList interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	Item(index int32) INode
}

type NodeList struct {
	length int32
}

func (n NodeList) GetLength() int32 {
	return n.length
}

func (n NodeList) SetLength(length int32) {
	n.length = length
}

func (n NodeList) Item(index int32) INode {
	return nil
}

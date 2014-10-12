package dom

type INamedNodeMap interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	GetNamedItem(name string) INode
	SetNamedItem(arg INode) INode
	RemoveNamedItem(name string) INode
	Item(index int32) INode
	GetNamedItemNS(namespaceURI string, localName string) INode
	SetNamedItemNS(arg INode) INode
	RemoveNamedItemNS(namespaceURI string, localName string) INode
}

type NamedNodeMap struct {
	length int32
}

func (n NamedNodeMap) GetLength() int32 {
	return n.length
}

func (n NamedNodeMap) SetLength(length int32) {
	n.length = length
}

func (n NamedNodeMap) GetNamedItem(name string) INode {
	return nil
}

func (n NamedNodeMap) SetNamedItem(arg INode) INode {
	return nil
}

func (n NamedNodeMap) RemoveNamedItem(name string) INode {
	return nil
}

func (n NamedNodeMap) Item(index int32) INode {
	return nil
}

func (n NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) INode {
	return nil
}

func (n NamedNodeMap) SetNamedItemNS(arg INode) INode {
	return nil
}

func (n NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) INode {
	return nil
}

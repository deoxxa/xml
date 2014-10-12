package dom

type NamedNodeMap interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	GetNamedItem(name string) Node
	SetNamedItem(arg Node) Node
	RemoveNamedItem(name string) Node
	Item(index int32) Node
	GetNamedItemNS(namespaceURI string, localName string) Node
	SetNamedItemNS(arg Node) Node
	RemoveNamedItemNS(namespaceURI string, localName string) Node
}

type NamedNodeMapImpl struct {
	length int32
}

func (n *NamedNodeMapImpl) GetLength() int32 {
	return n.length
}

func (n *NamedNodeMapImpl) SetLength(length int32) {
	n.length = length
}

func (n *NamedNodeMapImpl) GetNamedItem(name string) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NamedNodeMapImpl) SetNamedItem(arg Node) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NamedNodeMapImpl) RemoveNamedItem(name string) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NamedNodeMapImpl) Item(index int32) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NamedNodeMapImpl) GetNamedItemNS(namespaceURI string, localName string) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NamedNodeMapImpl) SetNamedItemNS(arg Node) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NamedNodeMapImpl) RemoveNamedItemNS(namespaceURI string, localName string) Node {
	panic("unimplemented") // TODO

	return nil
}

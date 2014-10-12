package dom

type INameList interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	GetName(index int32) string
	GetNamespaceURI(index int32) string
	Contains(str string) bool
	ContainsNS(namespaceURI string, name string) bool
}

type NameList struct {
	length int32
}

func (n *NameList) GetLength() int32 {
	return n.length
}

func (n *NameList) SetLength(length int32) {
	n.length = length
}

func (n *NameList) GetName(index int32) string {
	panic("unimplemented") // TODO

	return ""
}

func (n *NameList) GetNamespaceURI(index int32) string {
	panic("unimplemented") // TODO

	return ""
}

func (n *NameList) Contains(str string) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NameList) ContainsNS(namespaceURI string, name string) bool {
	panic("unimplemented") // TODO

	return false
}

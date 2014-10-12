package dom

type DOMImplementationList interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	Item(index int32) DOMImplementation
}

type DOMImplementationListImpl struct {
	length int32
}

func (d DOMImplementationListImpl) GetLength() int32 {
	return d.length
}

func (d DOMImplementationListImpl) SetLength(length int32) {
	d.length = length
}

func (d DOMImplementationListImpl) Item(index int32) DOMImplementation {
	panic("unimplemented") // TODO

	return nil
}

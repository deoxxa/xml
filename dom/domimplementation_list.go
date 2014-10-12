package dom

type IDOMImplementationList interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	Item(index int32) IDOMImplementation
}

type DOMImplementationList struct {
	length int32
}

func (d DOMImplementationList) GetLength() int32 {
	return d.length
}

func (d DOMImplementationList) SetLength(length int32) {
	d.length = length
}

func (d DOMImplementationList) Item(index int32) IDOMImplementation {
	panic("unimplemented") // TODO

	return nil
}

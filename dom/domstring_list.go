package dom

type IDOMStringList interface {
	// accessors
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	Item(index int32) string
	Contains(str string) bool
}

type DOMStringList struct {
	length int32
}

func (d DOMStringList) GetLength() int32 {
	return d.length
}

func (d DOMStringList) SetLength(length int32) {
	d.length = length
}

func (d DOMStringList) Item(index int32) string {
	panic("unimplemented") // TODO

	return ""
}

func (d DOMStringList) Contains(str string) bool {
	panic("unimplemented") // TODO

	return false
}

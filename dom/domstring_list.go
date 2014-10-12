package dom

type DOMStringList interface {
	// accessors
	GetLength() int
	// spec-defined functions
	Item(index int) string
	Contains(str string) bool
}

type DOMStringListImpl []string

func (d DOMStringListImpl) GetLength() int {
	return len(d)
}

func (d DOMStringListImpl) Item(index int) string {
	return d[index]
}

func (d DOMStringListImpl) Contains(str string) bool {
	for _, v := range d {
		if v == str {
			return true
		}
	}

	return false
}

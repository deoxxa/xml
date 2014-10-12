package dom

type IDOMStringList interface {
	// accessors
	GetLength() int
	// spec-defined functions
	Item(index int) string
	Contains(str string) bool
}

type DOMStringList []string

func (d DOMStringList) GetLength() int {
	return len(d)
}

func (d DOMStringList) Item(index int) string {
	return d[index]
}

func (d DOMStringList) Contains(str string) bool {
	for _, v := range d {
		if v == str {
			return true
		}
	}

	return false
}

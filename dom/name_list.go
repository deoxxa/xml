package dom

import (
	"strings"
)

type NameList interface {
	// accessors
	GetLength() int
	// spec-defined functions
	GetName(index int) string
	GetNamespaceURI(index int) string
	Contains(str string) bool
	ContainsNS(namespaceURI string, name string) bool
}

type NameListImpl []string

func (n NameListImpl) GetLength() int {
	return len(n)
}

func (n NameListImpl) GetName(index int) string {
	str := n[index]

	i := strings.Index(str, "}")

	return str[i:]
}

func (n NameListImpl) GetNamespaceURI(index int) string {
	str := n[index]

	i := strings.Index(str, "}")

	return str[1:i]
}

func (n NameListImpl) Contains(str string) bool {
	return n.ContainsNS("", str)
}

func (n NameListImpl) ContainsNS(namespaceURI string, name string) bool {
	str := "{" + namespaceURI + "}" + name

	for _, v := range n {
		if v == str {
			return true
		}
	}

	return false
}

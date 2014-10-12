package dom

import (
	"strings"
)

type INameList interface {
	// accessors
	GetLength() int
	// spec-defined functions
	GetName(index int) string
	GetNamespaceURI(index int) string
	Contains(str string) bool
	ContainsNS(namespaceURI string, name string) bool
}

type NameList []string

func (n NameList) GetLength() int {
	return len(n)
}

func (n NameList) GetName(index int) string {
	str := n[index]

	i := strings.Index(str, "}")

	return str[i:]
}

func (n NameList) GetNamespaceURI(index int) string {
	str := n[index]

	i := strings.Index(str, "}")

	return str[1:i]
}

func (n NameList) Contains(str string) bool {
	return n.ContainsNS("", str)
}

func (n NameList) ContainsNS(namespaceURI string, name string) bool {
	str := "{" + namespaceURI + "}" + name

	for _, v := range n {
		if v == str {
			return true
		}
	}

	return false
}

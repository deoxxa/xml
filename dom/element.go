package dom

type IElement interface {
	INode

	// accessors
	GetTagName() string
	SetTagName(tagName string)
	// spec-defined functions
	GetAttribute(name string) string
	SetAttribute(name string, value string)
	RemoveAttribute(name string)
	GetAttributeNode(name string) IAttr
	SetAttributeNode(newAttr IAttr) IAttr
	RemoveAttributeNode(oldAttr IAttr) IAttr
	GetElementsByTagName(name string) INodeList
	GetAttributeNS(namespaceURI string, localName string) string
	SetAttributeNS(namespaceURI string, qualifiedName string, value string)
	RemoveAttributeNS(namespaceURI string, localName string)
	GetAttributeNodeNS(namespaceURI string, localName string) IAttr
	SetAttributeNodeNS(newAttr IAttr) IAttr
	GetElementsByTagNameNS(namespaceURI string, localName string) INodeList
	HasAttribute(name string) bool
	HasAttributeNS(namespaceURI string, localName string) bool
}

type Element struct {
	Node
	tagName string
}

func (e Element) GetTagName() string {
	return e.tagName
}

func (e Element) SetTagName(tagName string) {
	e.tagName = tagName
}

func (e Element) GetAttribute(name string) string {
	return ""
}

func (e Element) SetAttribute(name string, value string) {
	return
}

func (e Element) RemoveAttribute(name string) {
	return
}

func (e Element) GetAttributeNode(name string) IAttr {
	return nil
}

func (e Element) SetAttributeNode(newAttr IAttr) IAttr {
	return nil
}

func (e Element) RemoveAttributeNode(oldAttr IAttr) IAttr {
	return nil
}

func (e Element) GetElementsByTagName(name string) INodeList {
	return nil
}

func (e Element) GetAttributeNS(namespaceURI string, localName string) string {
	return ""
}

func (e Element) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	return
}

func (e Element) RemoveAttributeNS(namespaceURI string, localName string) {
	return
}

func (e Element) GetAttributeNodeNS(namespaceURI string, localName string) IAttr {
	return nil
}

func (e Element) SetAttributeNodeNS(newAttr IAttr) IAttr {
	return nil
}

func (e Element) GetElementsByTagNameNS(namespaceURI string, localName string) INodeList {
	return nil
}

func (e Element) HasAttribute(name string) bool {
	return false
}

func (e Element) HasAttributeNS(namespaceURI string, localName string) bool {
	return false
}

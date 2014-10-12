package dom

type IElement interface {
	INode

	// accessors
	GetTagName() string
	SetTagName(tagName string)
	GetSchemaTypeInfo() ITypeInfo
	SetSchemaTypeInfo(schemaTypeInfo ITypeInfo)
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
	SetIdAttribute(name string, isId bool)
	SetIdAttributeNS(namespaceURI string, localName string, isId bool)
	SetIdAttributeNode(idAttr IAttr, isId bool)
}

type Element struct {
	Node
	tagName        string
	schemaTypeInfo ITypeInfo
}

func (e Element) GetTagName() string {
	return e.tagName
}

func (e Element) SetTagName(tagName string) {
	e.tagName = tagName
}

func (e Element) GetSchemaTypeInfo() ITypeInfo {
	return e.schemaTypeInfo
}

func (e Element) SetSchemaTypeInfo(schemaTypeInfo ITypeInfo) {
	e.schemaTypeInfo = schemaTypeInfo
}

func (e Element) GetAttribute(name string) string {
	panic("unimplemented") // TODO

	return ""
}

func (e Element) SetAttribute(name string, value string) {
	panic("unimplemented") // TODO

	return
}

func (e Element) RemoveAttribute(name string) {
	panic("unimplemented") // TODO

	return
}

func (e Element) GetAttributeNode(name string) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) SetAttributeNode(newAttr IAttr) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) RemoveAttributeNode(oldAttr IAttr) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) GetElementsByTagName(name string) INodeList {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) GetAttributeNS(namespaceURI string, localName string) string {
	panic("unimplemented") // TODO

	return ""
}

func (e Element) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	panic("unimplemented") // TODO

	return
}

func (e Element) RemoveAttributeNS(namespaceURI string, localName string) {
	panic("unimplemented") // TODO

	return
}

func (e Element) GetAttributeNodeNS(namespaceURI string, localName string) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) SetAttributeNodeNS(newAttr IAttr) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) GetElementsByTagNameNS(namespaceURI string, localName string) INodeList {
	panic("unimplemented") // TODO

	return nil
}

func (e Element) HasAttribute(name string) bool {
	panic("unimplemented") // TODO

	return false
}

func (e Element) HasAttributeNS(namespaceURI string, localName string) bool {
	panic("unimplemented") // TODO

	return false
}

func (e Element) SetIdAttribute(name string, isId bool) {
	panic("unimplemented") // TODO

	return
}

func (e Element) SetIdAttributeNS(namespaceURI string, localName string, isId bool) {
	panic("unimplemented") // TODO

	return
}

func (e Element) SetIdAttributeNode(idAttr IAttr, isId bool) {
	panic("unimplemented") // TODO

	return
}

package dom

type Element interface {
	Node

	// accessors
	GetTagName() string
	SetTagName(tagName string)
	GetSchemaTypeInfo() TypeInfo
	SetSchemaTypeInfo(schemaTypeInfo TypeInfo)
	// spec-defined functions
	GetAttribute(name string) string
	SetAttribute(name string, value string)
	RemoveAttribute(name string)
	GetAttributeNode(name string) Attr
	SetAttributeNode(newAttr Attr) Attr
	RemoveAttributeNode(oldAttr Attr) Attr
	GetElementsByTagName(name string) NodeList
	GetAttributeNS(namespaceURI string, localName string) string
	SetAttributeNS(namespaceURI string, qualifiedName string, value string)
	RemoveAttributeNS(namespaceURI string, localName string)
	GetAttributeNodeNS(namespaceURI string, localName string) Attr
	SetAttributeNodeNS(newAttr Attr) Attr
	GetElementsByTagNameNS(namespaceURI string, localName string) NodeList
	HasAttribute(name string) bool
	HasAttributeNS(namespaceURI string, localName string) bool
	SetIdAttribute(name string, isId bool)
	SetIdAttributeNS(namespaceURI string, localName string, isId bool)
	SetIdAttributeNode(idAttr Attr, isId bool)
}

type ElementImpl struct {
	NodeImpl
	tagName        string
	schemaTypeInfo TypeInfo
}

func (e *ElementImpl) GetTagName() string {
	return e.tagName
}

func (e *ElementImpl) SetTagName(tagName string) {
	e.tagName = tagName
}

func (e *ElementImpl) GetSchemaTypeInfo() TypeInfo {
	return e.schemaTypeInfo
}

func (e *ElementImpl) SetSchemaTypeInfo(schemaTypeInfo TypeInfo) {
	e.schemaTypeInfo = schemaTypeInfo
}

func (e *ElementImpl) GetAttribute(name string) string {
	panic("unimplemented") // TODO

	return ""
}

func (e *ElementImpl) SetAttribute(name string, value string) {
	panic("unimplemented") // TODO

	return
}

func (e *ElementImpl) RemoveAttribute(name string) {
	panic("unimplemented") // TODO

	return
}

func (e *ElementImpl) GetAttributeNode(name string) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) SetAttributeNode(newAttr Attr) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) RemoveAttributeNode(oldAttr Attr) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) GetElementsByTagName(name string) NodeList {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) GetAttributeNS(namespaceURI string, localName string) string {
	panic("unimplemented") // TODO

	return ""
}

func (e *ElementImpl) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	panic("unimplemented") // TODO

	return
}

func (e *ElementImpl) RemoveAttributeNS(namespaceURI string, localName string) {
	panic("unimplemented") // TODO

	return
}

func (e *ElementImpl) GetAttributeNodeNS(namespaceURI string, localName string) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) SetAttributeNodeNS(newAttr Attr) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) GetElementsByTagNameNS(namespaceURI string, localName string) NodeList {
	panic("unimplemented") // TODO

	return nil
}

func (e *ElementImpl) HasAttribute(name string) bool {
	panic("unimplemented") // TODO

	return false
}

func (e *ElementImpl) HasAttributeNS(namespaceURI string, localName string) bool {
	panic("unimplemented") // TODO

	return false
}

func (e *ElementImpl) SetIdAttribute(name string, isId bool) {
	panic("unimplemented") // TODO

	return
}

func (e *ElementImpl) SetIdAttributeNS(namespaceURI string, localName string, isId bool) {
	panic("unimplemented") // TODO

	return
}

func (e *ElementImpl) SetIdAttributeNode(idAttr Attr, isId bool) {
	panic("unimplemented") // TODO

	return
}

package dom

import (
	"strings"
)

type Document interface {
	Node

	// accessors
	GetDoctype() DocumentType
	SetDoctype(doctype DocumentType)
	GetImplementation() DOMImplementation
	SetImplementation(implementation DOMImplementation)
	GetDocumentElement() Element
	SetDocumentElement(documentElement Element)
	GetInputEncoding() string
	SetInputEncoding(inputEncoding string)
	GetXmlEncoding() string
	SetXmlEncoding(xmlEncoding string)
	GetXmlStandalone() bool
	SetXmlStandalone(xmlStandalone bool)
	GetXmlVersion() string
	SetXmlVersion(xmlVersion string)
	GetStrictErrorChecking() bool
	SetStrictErrorChecking(strictErrorChecking bool)
	GetDocumentURI() string
	SetDocumentURI(documentURI string)
	GetDomConfig() DOMConfiguration
	SetDomConfig(domConfig DOMConfiguration)
	// spec-defined functions
	CreateElement(tagName string) Element
	CreateDocumentFragment() DocumentFragment
	CreateTextNode(data string) Text
	CreateComment(data string) Comment
	CreateCDATASection(data string) CDATASection
	CreateProcessingInstruction(target string, data string) ProcessingInstruction
	CreateAttribute(name string) Attr
	CreateEntityReference(name string) EntityReference
	GetElementsByTagName(tagname string) NodeList
	ImportNode(importedNode Node, deep bool) Node
	CreateElementNS(namespaceURI string, qualifiedName string) Element
	CreateAttributeNS(namespaceURI string, qualifiedName string) Attr
	GetElementsByTagNameNS(namespaceURI string, localName string) NodeList
	GetElementById(elementId string) Element
	AdoptNode(source Node) Node
	NormalizeDocument()
	RenameNode(n Node, namespaceURI string, qualifiedName string) Node
}

type DocumentImpl struct {
	NodeImpl
	doctype             DocumentType
	implementation      DOMImplementation
	documentElement     Element
	inputEncoding       string
	xmlEncoding         string
	xmlStandalone       bool
	xmlVersion          string
	strictErrorChecking bool
	documentURI         string
	domConfig           DOMConfiguration
}

func (d *DocumentImpl) GetDoctype() DocumentType {
	return d.doctype
}

func (d *DocumentImpl) SetDoctype(doctype DocumentType) {
	d.doctype = doctype
}

func (d *DocumentImpl) GetImplementation() DOMImplementation {
	return d.implementation
}

func (d *DocumentImpl) SetImplementation(implementation DOMImplementation) {
	d.implementation = implementation
}

func (d *DocumentImpl) GetDocumentElement() Element {
	return d.documentElement
}

func (d *DocumentImpl) SetDocumentElement(documentElement Element) {
	d.documentElement = documentElement
}

func (d *DocumentImpl) GetInputEncoding() string {
	return d.inputEncoding
}

func (d *DocumentImpl) SetInputEncoding(inputEncoding string) {
	d.inputEncoding = inputEncoding
}

func (d *DocumentImpl) GetXmlEncoding() string {
	return d.xmlEncoding
}

func (d *DocumentImpl) SetXmlEncoding(xmlEncoding string) {
	d.xmlEncoding = xmlEncoding
}

func (d *DocumentImpl) GetXmlStandalone() bool {
	return d.xmlStandalone
}

func (d *DocumentImpl) SetXmlStandalone(xmlStandalone bool) {
	d.xmlStandalone = xmlStandalone
}

func (d *DocumentImpl) GetXmlVersion() string {
	return d.xmlVersion
}

func (d *DocumentImpl) SetXmlVersion(xmlVersion string) {
	d.xmlVersion = xmlVersion
}

func (d *DocumentImpl) GetStrictErrorChecking() bool {
	return d.strictErrorChecking
}

func (d *DocumentImpl) SetStrictErrorChecking(strictErrorChecking bool) {
	d.strictErrorChecking = strictErrorChecking
}

func (d *DocumentImpl) GetDocumentURI() string {
	return d.documentURI
}

func (d *DocumentImpl) SetDocumentURI(documentURI string) {
	d.documentURI = documentURI
}

func (d *DocumentImpl) GetDomConfig() DOMConfiguration {
	return d.domConfig
}

func (d *DocumentImpl) SetDomConfig(domConfig DOMConfiguration) {
	d.domConfig = domConfig
}

func (d *DocumentImpl) InsertBefore(newChild Node, refChild Node) Node {
	if d.documentElement == nil {
		d.documentElement = newChild.(*ElementImpl)

		return d.NodeImpl.InsertBefore(newChild, refChild)
	} else {
		return d.documentElement.InsertBefore(newChild, refChild)
	}
}

func (d *DocumentImpl) AppendChild(newChild Node) Node {
	return d.InsertBefore(newChild, nil)
}

func (d *DocumentImpl) CreateElement(tagName string) Element {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateDocumentFragment() DocumentFragment {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateTextNode(data string) Text {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateComment(data string) Comment {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateCDATASection(data string) CDATASection {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateProcessingInstruction(target string, data string) ProcessingInstruction {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateAttribute(name string) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateEntityReference(name string) EntityReference {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) GetElementsByTagName(tagname string) NodeList {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) ImportNode(importedNode Node, deep bool) Node {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) CreateElementNS(namespaceURI string, qualifiedName string) Element {
	index := strings.LastIndex(qualifiedName, ":")

	var prefix, localName string
	if index != -1 {
		prefix = qualifiedName[0:index]
		localName = qualifiedName[index+1:]
	} else {
		prefix = ""
		localName = qualifiedName
	}

	return &ElementImpl{
		NodeImpl: NodeImpl{
			nodeName:      qualifiedName,
			childNodes:    NodeListImpl{},
			attributes:    &NamedNodeMapImpl{},
			ownerDocument: d,
			namespaceURI:  namespaceURI,
			prefix:        prefix,
			localName:     localName,
		},
		tagName: qualifiedName,
	}
}

func (d *DocumentImpl) CreateAttributeNS(namespaceURI string, qualifiedName string) Attr {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) GetElementsByTagNameNS(namespaceURI string, localName string) NodeList {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) GetElementById(elementId string) Element {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) AdoptNode(source Node) Node {
	panic("unimplemented") // TODO

	return nil
}

func (d *DocumentImpl) NormalizeDocument() {
	panic("unimplemented") // TODO

	return
}

func (d *DocumentImpl) RenameNode(n Node, namespaceURI string, qualifiedName string) Node {
	panic("unimplemented") // TODO

	return nil
}

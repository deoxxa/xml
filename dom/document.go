package dom

type IDocument interface {
	INode

	// accessors
	GetDoctype() IDocumentType
	SetDoctype(doctype IDocumentType)
	GetImplementation() IDOMImplementation
	SetImplementation(implementation IDOMImplementation)
	GetDocumentElement() IElement
	SetDocumentElement(documentElement IElement)
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
	GetDomConfig() IDOMConfiguration
	SetDomConfig(domConfig IDOMConfiguration)
	// spec-defined functions
	CreateElement(tagName string) IElement
	CreateDocumentFragment() IDocumentFragment
	CreateTextNode(data string) IText
	CreateComment(data string) IComment
	CreateCDATASection(data string) ICDATASection
	CreateProcessingInstruction(target string, data string) IProcessingInstruction
	CreateAttribute(name string) IAttr
	CreateEntityReference(name string) IEntityReference
	GetElementsByTagName(tagname string) INodeList
	ImportNode(importedNode INode, deep bool) INode
	CreateElementNS(namespaceURI string, qualifiedName string) IElement
	CreateAttributeNS(namespaceURI string, qualifiedName string) IAttr
	GetElementsByTagNameNS(namespaceURI string, localName string) INodeList
	GetElementById(elementId string) IElement
	AdoptNode(source INode) INode
	NormalizeDocument()
	RenameNode(n INode, namespaceURI string, qualifiedName string) INode
}

type Document struct {
	Node
	doctype             IDocumentType
	implementation      IDOMImplementation
	documentElement     IElement
	inputEncoding       string
	xmlEncoding         string
	xmlStandalone       bool
	xmlVersion          string
	strictErrorChecking bool
	documentURI         string
	domConfig           IDOMConfiguration
}

func (d Document) GetDoctype() IDocumentType {
	return d.doctype
}

func (d Document) SetDoctype(doctype IDocumentType) {
	d.doctype = doctype
}

func (d Document) GetImplementation() IDOMImplementation {
	return d.implementation
}

func (d Document) SetImplementation(implementation IDOMImplementation) {
	d.implementation = implementation
}

func (d Document) GetDocumentElement() IElement {
	return d.documentElement
}

func (d Document) SetDocumentElement(documentElement IElement) {
	d.documentElement = documentElement
}

func (d Document) GetInputEncoding() string {
	return d.inputEncoding
}

func (d Document) SetInputEncoding(inputEncoding string) {
	d.inputEncoding = inputEncoding
}

func (d Document) GetXmlEncoding() string {
	return d.xmlEncoding
}

func (d Document) SetXmlEncoding(xmlEncoding string) {
	d.xmlEncoding = xmlEncoding
}

func (d Document) GetXmlStandalone() bool {
	return d.xmlStandalone
}

func (d Document) SetXmlStandalone(xmlStandalone bool) {
	d.xmlStandalone = xmlStandalone
}

func (d Document) GetXmlVersion() string {
	return d.xmlVersion
}

func (d Document) SetXmlVersion(xmlVersion string) {
	d.xmlVersion = xmlVersion
}

func (d Document) GetStrictErrorChecking() bool {
	return d.strictErrorChecking
}

func (d Document) SetStrictErrorChecking(strictErrorChecking bool) {
	d.strictErrorChecking = strictErrorChecking
}

func (d Document) GetDocumentURI() string {
	return d.documentURI
}

func (d Document) SetDocumentURI(documentURI string) {
	d.documentURI = documentURI
}

func (d Document) GetDomConfig() IDOMConfiguration {
	return d.domConfig
}

func (d Document) SetDomConfig(domConfig IDOMConfiguration) {
	d.domConfig = domConfig
}

func (d Document) CreateElement(tagName string) IElement {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateDocumentFragment() IDocumentFragment {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateTextNode(data string) IText {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateComment(data string) IComment {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateCDATASection(data string) ICDATASection {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateProcessingInstruction(target string, data string) IProcessingInstruction {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateAttribute(name string) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateEntityReference(name string) IEntityReference {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) GetElementsByTagName(tagname string) INodeList {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) ImportNode(importedNode INode, deep bool) INode {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateElementNS(namespaceURI string, qualifiedName string) IElement {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) CreateAttributeNS(namespaceURI string, qualifiedName string) IAttr {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) GetElementsByTagNameNS(namespaceURI string, localName string) INodeList {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) GetElementById(elementId string) IElement {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) AdoptNode(source INode) INode {
	panic("unimplemented") // TODO

	return nil
}

func (d Document) NormalizeDocument() {
	panic("unimplemented") // TODO

	return
}

func (d Document) RenameNode(n INode, namespaceURI string, qualifiedName string) INode {
	panic("unimplemented") // TODO

	return nil
}

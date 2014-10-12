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
}

type Document struct {
	Node
	doctype         IDocumentType
	implementation  IDOMImplementation
	documentElement IElement
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

func (d Document) CreateElement(tagName string) IElement {
	return nil
}

func (d Document) CreateDocumentFragment() IDocumentFragment {
	return nil
}

func (d Document) CreateTextNode(data string) IText {
	return nil
}

func (d Document) CreateComment(data string) IComment {
	return nil
}

func (d Document) CreateCDATASection(data string) ICDATASection {
	return nil
}

func (d Document) CreateProcessingInstruction(target string, data string) IProcessingInstruction {
	return nil
}

func (d Document) CreateAttribute(name string) IAttr {
	return nil
}

func (d Document) CreateEntityReference(name string) IEntityReference {
	return nil
}

func (d Document) GetElementsByTagName(tagname string) INodeList {
	return nil
}

func (d Document) ImportNode(importedNode INode, deep bool) INode {
	return nil
}

func (d Document) CreateElementNS(namespaceURI string, qualifiedName string) IElement {
	return nil
}

func (d Document) CreateAttributeNS(namespaceURI string, qualifiedName string) IAttr {
	return nil
}

func (d Document) GetElementsByTagNameNS(namespaceURI string, localName string) INodeList {
	return nil
}

func (d Document) GetElementById(elementId string) IElement {
	return nil
}

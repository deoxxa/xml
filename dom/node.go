package dom

type Node interface {
	// accessors
	GetNodeName() string
	SetNodeName(nodeName string)
	GetNodeValue() string
	SetNodeValue(nodeValue string)
	GetNodeType() int16
	SetNodeType(nodeType int16)
	GetParentNode() Node
	SetParentNode(parentNode Node)
	GetChildNodes() NodeList
	SetChildNodes(childNodes NodeList)
	GetFirstChild() Node
	SetFirstChild(firstChild Node)
	GetLastChild() Node
	SetLastChild(lastChild Node)
	GetPreviousSibling() Node
	SetPreviousSibling(previousSibling Node)
	GetNextSibling() Node
	SetNextSibling(nextSibling Node)
	GetAttributes() NamedNodeMap
	SetAttributes(attributes NamedNodeMap)
	GetOwnerDocument() Document
	SetOwnerDocument(ownerDocument Document)
	GetNamespaceURI() string
	SetNamespaceURI(namespaceURI string)
	GetPrefix() string
	SetPrefix(prefix string)
	GetLocalName() string
	SetLocalName(localName string)
	GetBaseURI() string
	SetBaseURI(baseURI string)
	GetTextContent() string
	SetTextContent(textContent string)
	// spec-defined functions
	InsertBefore(newChild Node, refChild Node) Node
	ReplaceChild(newChild Node, oldChild Node) Node
	RemoveChild(oldChild Node) Node
	AppendChild(newChild Node) Node
	HasChildNodes() bool
	CloneNode(deep bool) Node
	Normalize()
	IsSupported(feature string, version string) bool
	HasAttributes() bool
	CompareDocumentPosition(other Node) int16
	IsSameNode(other Node) bool
	LookupPrefix(namespaceURI string) string
	IsDefaultNamespace(namespaceURI string) bool
	LookupNamespaceURI(prefix string) string
	IsEqualNode(arg Node) bool
	GetFeature(feature string, version string) DOMObject
	SetUserData(key string, data DOMUserData, handler UserDataHandler) DOMUserData
	GetUserData(key string) DOMUserData
}

const (
	ELEMENT_NODE                              = 1
	ATTRIBUTE_NODE                            = 2
	TEXT_NODE                                 = 3
	CDATA_SECTION_NODE                        = 4
	ENTITY_REFERENCE_NODE                     = 5
	ENTITY_NODE                               = 6
	PROCESSING_INSTRUCTION_NODE               = 7
	COMMENT_NODE                              = 8
	DOCUMENT_NODE                             = 9
	DOCUMENT_TYPE_NODE                        = 10
	DOCUMENT_FRAGMENT_NODE                    = 11
	NOTATION_NODE                             = 12
	DOCUMENT_POSITION_DISCONNECTED            = 0x01
	DOCUMENT_POSITION_PRECEDING               = 0x02
	DOCUMENT_POSITION_FOLLOWING               = 0x04
	DOCUMENT_POSITION_CONTAINS                = 0x08
	DOCUMENT_POSITION_CONTAINED_BY            = 0x10
	DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20
)

type NodeImpl struct {
	nodeName        string
	nodeValue       string
	nodeType        int16
	parentNode      Node
	childNodes      NodeList
	firstChild      Node
	lastChild       Node
	previousSibling Node
	nextSibling     Node
	attributes      NamedNodeMap
	ownerDocument   Document
	namespaceURI    string
	prefix          string
	localName       string
	baseURI         string
	textContent     string
}

func (n *NodeImpl) GetNodeName() string {
	return n.nodeName
}

func (n *NodeImpl) SetNodeName(nodeName string) {
	n.nodeName = nodeName
}

func (n *NodeImpl) GetNodeValue() string {
	return n.nodeValue
}

func (n *NodeImpl) SetNodeValue(nodeValue string) {
	n.nodeValue = nodeValue
}

func (n *NodeImpl) GetNodeType() int16 {
	return n.nodeType
}

func (n *NodeImpl) SetNodeType(nodeType int16) {
	n.nodeType = nodeType
}

func (n *NodeImpl) GetParentNode() Node {
	return n.parentNode
}

func (n *NodeImpl) SetParentNode(parentNode Node) {
	n.parentNode = parentNode
}

func (n *NodeImpl) GetChildNodes() NodeList {
	return n.childNodes
}

func (n *NodeImpl) SetChildNodes(childNodes NodeList) {
	n.childNodes = childNodes
}

func (n *NodeImpl) GetFirstChild() Node {
	return n.firstChild
}

func (n *NodeImpl) SetFirstChild(firstChild Node) {
	n.firstChild = firstChild
}

func (n *NodeImpl) GetLastChild() Node {
	return n.lastChild
}

func (n *NodeImpl) SetLastChild(lastChild Node) {
	n.lastChild = lastChild
}

func (n *NodeImpl) GetPreviousSibling() Node {
	return n.previousSibling
}

func (n *NodeImpl) SetPreviousSibling(previousSibling Node) {
	n.previousSibling = previousSibling
}

func (n *NodeImpl) GetNextSibling() Node {
	return n.nextSibling
}

func (n *NodeImpl) SetNextSibling(nextSibling Node) {
	n.nextSibling = nextSibling
}

func (n *NodeImpl) GetAttributes() NamedNodeMap {
	return n.attributes
}

func (n *NodeImpl) SetAttributes(attributes NamedNodeMap) {
	n.attributes = attributes
}

func (n *NodeImpl) GetOwnerDocument() Document {
	return n.ownerDocument
}

func (n *NodeImpl) SetOwnerDocument(ownerDocument Document) {
	n.ownerDocument = ownerDocument
}

func (n *NodeImpl) GetNamespaceURI() string {
	return n.namespaceURI
}

func (n *NodeImpl) SetNamespaceURI(namespaceURI string) {
	n.namespaceURI = namespaceURI
}

func (n *NodeImpl) GetPrefix() string {
	return n.prefix
}

func (n *NodeImpl) SetPrefix(prefix string) {
	n.prefix = prefix
}

func (n *NodeImpl) GetLocalName() string {
	return n.localName
}

func (n *NodeImpl) SetLocalName(localName string) {
	n.localName = localName
}

func (n *NodeImpl) GetBaseURI() string {
	return n.baseURI
}

func (n *NodeImpl) SetBaseURI(baseURI string) {
	n.baseURI = baseURI
}

func (n *NodeImpl) GetTextContent() string {
	return n.textContent
}

func (n *NodeImpl) SetTextContent(textContent string) {
	n.textContent = textContent
}

func (n *NodeImpl) InsertBefore(newChild Node, refChild Node) Node {
	if newChild.GetParentNode() != nil {
		newChild.GetParentNode().RemoveChild(newChild)
	}

	var previousSibling Node
	if refChild != nil {
		previousSibling = refChild.GetPreviousSibling()
	} else {
		previousSibling = n.GetLastChild()
	}

	newChild.SetPreviousSibling(previousSibling)
	newChild.SetNextSibling(refChild)

	if previousSibling != nil {
		previousSibling.SetNextSibling(newChild)
	} else {
		n.firstChild = newChild
	}

	if refChild != nil {
		refChild.SetPreviousSibling(newChild)
	} else {
		n.lastChild = newChild
	}

	newChild.SetParentNode(n)

	n.childNodes = append(n.childNodes.(NodeListImpl), newChild)

	return newChild
}

func (n *NodeImpl) ReplaceChild(newChild Node, oldChild Node) Node {
	n.InsertBefore(newChild, oldChild)

	if oldChild != nil {
		n.RemoveChild(oldChild)
	}

	return newChild
}

func (n *NodeImpl) RemoveChild(oldChild Node) Node {
	previousSibling := oldChild.GetPreviousSibling()
	nextSibling := oldChild.GetNextSibling()

	if previousSibling != nil {
		previousSibling.SetNextSibling(nextSibling)
	} else {
		n.firstChild = nextSibling
	}

	if nextSibling != nil {
		nextSibling.SetPreviousSibling(previousSibling)
	} else {
		n.lastChild = previousSibling
	}

	return oldChild
}

func (n *NodeImpl) AppendChild(newChild Node) Node {
	return n.InsertBefore(newChild, nil)
}

func (n *NodeImpl) HasChildNodes() bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NodeImpl) CloneNode(deep bool) Node {
	panic("unimplemented") // TODO

	return nil
}

func (n *NodeImpl) Normalize() {
	panic("unimplemented") // TODO

	return
}

func (n *NodeImpl) IsSupported(feature string, version string) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NodeImpl) HasAttributes() bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NodeImpl) CompareDocumentPosition(other Node) int16 {
	panic("unimplemented") // TODO

	return 0
}

func (n *NodeImpl) IsSameNode(other Node) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NodeImpl) LookupPrefix(namespaceURI string) string {
	panic("unimplemented") // TODO

	return ""
}

func (n *NodeImpl) IsDefaultNamespace(namespaceURI string) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NodeImpl) LookupNamespaceURI(prefix string) string {
	panic("unimplemented") // TODO

	return ""
}

func (n *NodeImpl) IsEqualNode(arg Node) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *NodeImpl) GetFeature(feature string, version string) DOMObject {
	panic("unimplemented") // TODO

	return nil
}

func (n *NodeImpl) SetUserData(key string, data DOMUserData, handler UserDataHandler) DOMUserData {
	panic("unimplemented") // TODO

	return nil
}

func (n *NodeImpl) GetUserData(key string) DOMUserData {
	panic("unimplemented") // TODO

	return nil
}

package dom

type INode interface {
	// accessors
	GetNodeName() string
	SetNodeName(nodeName string)
	GetNodeValue() string
	SetNodeValue(nodeValue string)
	GetNodeType() int16
	SetNodeType(nodeType int16)
	GetParentNode() INode
	SetParentNode(parentNode INode)
	GetChildNodes() INodeList
	SetChildNodes(childNodes INodeList)
	GetFirstChild() INode
	SetFirstChild(firstChild INode)
	GetLastChild() INode
	SetLastChild(lastChild INode)
	GetPreviousSibling() INode
	SetPreviousSibling(previousSibling INode)
	GetNextSibling() INode
	SetNextSibling(nextSibling INode)
	GetAttributes() INamedNodeMap
	SetAttributes(attributes INamedNodeMap)
	GetOwnerDocument() IDocument
	SetOwnerDocument(ownerDocument IDocument)
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
	InsertBefore(newChild INode, refChild INode) INode
	ReplaceChild(newChild INode, oldChild INode) INode
	RemoveChild(oldChild INode) INode
	AppendChild(newChild INode) INode
	HasChildNodes() bool
	CloneNode(deep bool) INode
	Normalize()
	IsSupported(feature string, version string) bool
	HasAttributes() bool
	CompareDocumentPosition(other INode) int16
	IsSameNode(other INode) bool
	LookupPrefix(namespaceURI string) string
	IsDefaultNamespace(namespaceURI string) bool
	LookupNamespaceURI(prefix string) string
	IsEqualNode(arg INode) bool
	GetFeature(feature string, version string) IDOMObject
	SetUserData(key string, data IDOMUserData, handler IUserDataHandler) IDOMUserData
	GetUserData(key string) IDOMUserData
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

type Node struct {
	nodeName        string
	nodeValue       string
	nodeType        int16
	parentNode      INode
	childNodes      INodeList
	firstChild      INode
	lastChild       INode
	previousSibling INode
	nextSibling     INode
	attributes      INamedNodeMap
	ownerDocument   IDocument
	namespaceURI    string
	prefix          string
	localName       string
	baseURI         string
	textContent     string
}

func (n *Node) GetNodeName() string {
	return n.nodeName
}

func (n *Node) SetNodeName(nodeName string) {
	n.nodeName = nodeName
}

func (n *Node) GetNodeValue() string {
	return n.nodeValue
}

func (n *Node) SetNodeValue(nodeValue string) {
	n.nodeValue = nodeValue
}

func (n *Node) GetNodeType() int16 {
	return n.nodeType
}

func (n *Node) SetNodeType(nodeType int16) {
	n.nodeType = nodeType
}

func (n *Node) GetParentNode() INode {
	return n.parentNode
}

func (n *Node) SetParentNode(parentNode INode) {
	n.parentNode = parentNode
}

func (n *Node) GetChildNodes() INodeList {
	return n.childNodes
}

func (n *Node) SetChildNodes(childNodes INodeList) {
	n.childNodes = childNodes
}

func (n *Node) GetFirstChild() INode {
	return n.firstChild
}

func (n *Node) SetFirstChild(firstChild INode) {
	n.firstChild = firstChild
}

func (n *Node) GetLastChild() INode {
	return n.lastChild
}

func (n *Node) SetLastChild(lastChild INode) {
	n.lastChild = lastChild
}

func (n *Node) GetPreviousSibling() INode {
	return n.previousSibling
}

func (n *Node) SetPreviousSibling(previousSibling INode) {
	n.previousSibling = previousSibling
}

func (n *Node) GetNextSibling() INode {
	return n.nextSibling
}

func (n *Node) SetNextSibling(nextSibling INode) {
	n.nextSibling = nextSibling
}

func (n *Node) GetAttributes() INamedNodeMap {
	return n.attributes
}

func (n *Node) SetAttributes(attributes INamedNodeMap) {
	n.attributes = attributes
}

func (n *Node) GetOwnerDocument() IDocument {
	return n.ownerDocument
}

func (n *Node) SetOwnerDocument(ownerDocument IDocument) {
	n.ownerDocument = ownerDocument
}

func (n *Node) GetNamespaceURI() string {
	return n.namespaceURI
}

func (n *Node) SetNamespaceURI(namespaceURI string) {
	n.namespaceURI = namespaceURI
}

func (n *Node) GetPrefix() string {
	return n.prefix
}

func (n *Node) SetPrefix(prefix string) {
	n.prefix = prefix
}

func (n *Node) GetLocalName() string {
	return n.localName
}

func (n *Node) SetLocalName(localName string) {
	n.localName = localName
}

func (n *Node) GetBaseURI() string {
	return n.baseURI
}

func (n *Node) SetBaseURI(baseURI string) {
	n.baseURI = baseURI
}

func (n *Node) GetTextContent() string {
	return n.textContent
}

func (n *Node) SetTextContent(textContent string) {
	n.textContent = textContent
}

func (n *Node) InsertBefore(newChild INode, refChild INode) INode {
	if newChild.GetParentNode() != nil {
		newChild.GetParentNode().RemoveChild(newChild)
	}

	var previousSibling INode
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

	n.childNodes = append(n.childNodes.(NodeList), newChild)

	return newChild
}

func (n *Node) ReplaceChild(newChild INode, oldChild INode) INode {
	n.InsertBefore(newChild, oldChild)

	if oldChild != nil {
		n.RemoveChild(oldChild)
	}

	return newChild
}

func (n *Node) RemoveChild(oldChild INode) INode {
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

func (n *Node) AppendChild(newChild INode) INode {
	return n.InsertBefore(newChild, nil)
}

func (n *Node) HasChildNodes() bool {
	panic("unimplemented") // TODO

	return false
}

func (n *Node) CloneNode(deep bool) INode {
	panic("unimplemented") // TODO

	return nil
}

func (n *Node) Normalize() {
	panic("unimplemented") // TODO

	return
}

func (n *Node) IsSupported(feature string, version string) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *Node) HasAttributes() bool {
	panic("unimplemented") // TODO

	return false
}

func (n *Node) CompareDocumentPosition(other INode) int16 {
	panic("unimplemented") // TODO

	return 0
}

func (n *Node) IsSameNode(other INode) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *Node) LookupPrefix(namespaceURI string) string {
	panic("unimplemented") // TODO

	return ""
}

func (n *Node) IsDefaultNamespace(namespaceURI string) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *Node) LookupNamespaceURI(prefix string) string {
	panic("unimplemented") // TODO

	return ""
}

func (n *Node) IsEqualNode(arg INode) bool {
	panic("unimplemented") // TODO

	return false
}

func (n *Node) GetFeature(feature string, version string) IDOMObject {
	panic("unimplemented") // TODO

	return nil
}

func (n *Node) SetUserData(key string, data IDOMUserData, handler IUserDataHandler) IDOMUserData {
	panic("unimplemented") // TODO

	return nil
}

func (n *Node) GetUserData(key string) IDOMUserData {
	panic("unimplemented") // TODO

	return nil
}

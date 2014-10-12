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
}

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
}

func (n Node) GetNodeName() string {
	return n.nodeName
}

func (n Node) SetNodeName(nodeName string) {
	n.nodeName = nodeName
}

func (n Node) GetNodeValue() string {
	return n.nodeValue
}

func (n Node) SetNodeValue(nodeValue string) {
	n.nodeValue = nodeValue
}

func (n Node) GetNodeType() int16 {
	return n.nodeType
}

func (n Node) SetNodeType(nodeType int16) {
	n.nodeType = nodeType
}

func (n Node) GetParentNode() INode {
	return n.parentNode
}

func (n Node) SetParentNode(parentNode INode) {
	n.parentNode = parentNode
}

func (n Node) GetChildNodes() INodeList {
	return n.childNodes
}

func (n Node) SetChildNodes(childNodes INodeList) {
	n.childNodes = childNodes
}

func (n Node) GetFirstChild() INode {
	return n.firstChild
}

func (n Node) SetFirstChild(firstChild INode) {
	n.firstChild = firstChild
}

func (n Node) GetLastChild() INode {
	return n.lastChild
}

func (n Node) SetLastChild(lastChild INode) {
	n.lastChild = lastChild
}

func (n Node) GetPreviousSibling() INode {
	return n.previousSibling
}

func (n Node) SetPreviousSibling(previousSibling INode) {
	n.previousSibling = previousSibling
}

func (n Node) GetNextSibling() INode {
	return n.nextSibling
}

func (n Node) SetNextSibling(nextSibling INode) {
	n.nextSibling = nextSibling
}

func (n Node) GetAttributes() INamedNodeMap {
	return n.attributes
}

func (n Node) SetAttributes(attributes INamedNodeMap) {
	n.attributes = attributes
}

func (n Node) GetOwnerDocument() IDocument {
	return n.ownerDocument
}

func (n Node) SetOwnerDocument(ownerDocument IDocument) {
	n.ownerDocument = ownerDocument
}

func (n Node) GetNamespaceURI() string {
	return n.namespaceURI
}

func (n Node) SetNamespaceURI(namespaceURI string) {
	n.namespaceURI = namespaceURI
}

func (n Node) GetPrefix() string {
	return n.prefix
}

func (n Node) SetPrefix(prefix string) {
	n.prefix = prefix
}

func (n Node) GetLocalName() string {
	return n.localName
}

func (n Node) SetLocalName(localName string) {
	n.localName = localName
}

func (n Node) InsertBefore(newChild INode, refChild INode) INode {
	return nil
}

func (n Node) ReplaceChild(newChild INode, oldChild INode) INode {
	return nil
}

func (n Node) RemoveChild(oldChild INode) INode {
	return nil
}

func (n Node) AppendChild(newChild INode) INode {
	return nil
}

func (n Node) HasChildNodes() bool {
	return false
}

func (n Node) CloneNode(deep bool) INode {
	return nil
}

func (n Node) Normalize() {
	return
}

func (n Node) IsSupported(feature string, version string) bool {
	return false
}

func (n Node) HasAttributes() bool {
	return false
}

package dom

type IDOMLocator interface {
	// accessors
	GetLineNumber() int32
	SetLineNumber(lineNumber int32)
	GetColumnNumber() int32
	SetColumnNumber(columnNumber int32)
	GetByteOffset() int32
	SetByteOffset(byteOffset int32)
	GetUtf16Offset() int32
	SetUtf16Offset(utf16Offset int32)
	GetRelatedNode() INode
	SetRelatedNode(relatedNode INode)
	GetUri() string
	SetUri(uri string)
	// no spec-defined functions
}

type DOMLocator struct {
	lineNumber   int32
	columnNumber int32
	byteOffset   int32
	utf16Offset  int32
	relatedNode  INode
	uri          string
}

func (d *DOMLocator) GetLineNumber() int32 {
	return d.lineNumber
}

func (d *DOMLocator) SetLineNumber(lineNumber int32) {
	d.lineNumber = lineNumber
}

func (d *DOMLocator) GetColumnNumber() int32 {
	return d.columnNumber
}

func (d *DOMLocator) SetColumnNumber(columnNumber int32) {
	d.columnNumber = columnNumber
}

func (d *DOMLocator) GetByteOffset() int32 {
	return d.byteOffset
}

func (d *DOMLocator) SetByteOffset(byteOffset int32) {
	d.byteOffset = byteOffset
}

func (d *DOMLocator) GetUtf16Offset() int32 {
	return d.utf16Offset
}

func (d *DOMLocator) SetUtf16Offset(utf16Offset int32) {
	d.utf16Offset = utf16Offset
}

func (d *DOMLocator) GetRelatedNode() INode {
	return d.relatedNode
}

func (d *DOMLocator) SetRelatedNode(relatedNode INode) {
	d.relatedNode = relatedNode
}

func (d *DOMLocator) GetUri() string {
	return d.uri
}

func (d *DOMLocator) SetUri(uri string) {
	d.uri = uri
}

package dom

type DOMLocator interface {
	// accessors
	GetLineNumber() int32
	SetLineNumber(lineNumber int32)
	GetColumnNumber() int32
	SetColumnNumber(columnNumber int32)
	GetByteOffset() int32
	SetByteOffset(byteOffset int32)
	GetUtf16Offset() int32
	SetUtf16Offset(utf16Offset int32)
	GetRelatedNode() Node
	SetRelatedNode(relatedNode Node)
	GetUri() string
	SetUri(uri string)
	// no spec-defined functions
}

type DOMLocatorImpl struct {
	lineNumber   int32
	columnNumber int32
	byteOffset   int32
	utf16Offset  int32
	relatedNode  Node
	uri          string
}

func (d *DOMLocatorImpl) GetLineNumber() int32 {
	return d.lineNumber
}

func (d *DOMLocatorImpl) SetLineNumber(lineNumber int32) {
	d.lineNumber = lineNumber
}

func (d *DOMLocatorImpl) GetColumnNumber() int32 {
	return d.columnNumber
}

func (d *DOMLocatorImpl) SetColumnNumber(columnNumber int32) {
	d.columnNumber = columnNumber
}

func (d *DOMLocatorImpl) GetByteOffset() int32 {
	return d.byteOffset
}

func (d *DOMLocatorImpl) SetByteOffset(byteOffset int32) {
	d.byteOffset = byteOffset
}

func (d *DOMLocatorImpl) GetUtf16Offset() int32 {
	return d.utf16Offset
}

func (d *DOMLocatorImpl) SetUtf16Offset(utf16Offset int32) {
	d.utf16Offset = utf16Offset
}

func (d *DOMLocatorImpl) GetRelatedNode() Node {
	return d.relatedNode
}

func (d *DOMLocatorImpl) SetRelatedNode(relatedNode Node) {
	d.relatedNode = relatedNode
}

func (d *DOMLocatorImpl) GetUri() string {
	return d.uri
}

func (d *DOMLocatorImpl) SetUri(uri string) {
	d.uri = uri
}

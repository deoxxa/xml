package dom

type Entity interface {
	Node

	// accessors
	GetPublicId() string
	SetPublicId(publicId string)
	GetSystemId() string
	SetSystemId(systemId string)
	GetNotationName() string
	SetNotationName(notationName string)
	GetInputEncoding() string
	SetInputEncoding(inputEncoding string)
	GetXmlEncoding() string
	SetXmlEncoding(xmlEncoding string)
	GetXmlVersion() string
	SetXmlVersion(xmlVersion string)
	// no spec-defined functions
}

type EntityImpl struct {
	NodeImpl
	publicId      string
	systemId      string
	notationName  string
	inputEncoding string
	xmlEncoding   string
	xmlVersion    string
}

func (e *EntityImpl) GetPublicId() string {
	return e.publicId
}

func (e *EntityImpl) SetPublicId(publicId string) {
	e.publicId = publicId
}

func (e *EntityImpl) GetSystemId() string {
	return e.systemId
}

func (e *EntityImpl) SetSystemId(systemId string) {
	e.systemId = systemId
}

func (e *EntityImpl) GetNotationName() string {
	return e.notationName
}

func (e *EntityImpl) SetNotationName(notationName string) {
	e.notationName = notationName
}

func (e *EntityImpl) GetInputEncoding() string {
	return e.inputEncoding
}

func (e *EntityImpl) SetInputEncoding(inputEncoding string) {
	e.inputEncoding = inputEncoding
}

func (e *EntityImpl) GetXmlEncoding() string {
	return e.xmlEncoding
}

func (e *EntityImpl) SetXmlEncoding(xmlEncoding string) {
	e.xmlEncoding = xmlEncoding
}

func (e *EntityImpl) GetXmlVersion() string {
	return e.xmlVersion
}

func (e *EntityImpl) SetXmlVersion(xmlVersion string) {
	e.xmlVersion = xmlVersion
}

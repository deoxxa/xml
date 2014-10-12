package dom

type IEntity interface {
	INode

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

type Entity struct {
	Node
	publicId      string
	systemId      string
	notationName  string
	inputEncoding string
	xmlEncoding   string
	xmlVersion    string
}

func (e Entity) GetPublicId() string {
	return e.publicId
}

func (e Entity) SetPublicId(publicId string) {
	e.publicId = publicId
}

func (e Entity) GetSystemId() string {
	return e.systemId
}

func (e Entity) SetSystemId(systemId string) {
	e.systemId = systemId
}

func (e Entity) GetNotationName() string {
	return e.notationName
}

func (e Entity) SetNotationName(notationName string) {
	e.notationName = notationName
}

func (e Entity) GetInputEncoding() string {
	return e.inputEncoding
}

func (e Entity) SetInputEncoding(inputEncoding string) {
	e.inputEncoding = inputEncoding
}

func (e Entity) GetXmlEncoding() string {
	return e.xmlEncoding
}

func (e Entity) SetXmlEncoding(xmlEncoding string) {
	e.xmlEncoding = xmlEncoding
}

func (e Entity) GetXmlVersion() string {
	return e.xmlVersion
}

func (e Entity) SetXmlVersion(xmlVersion string) {
	e.xmlVersion = xmlVersion
}

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
	// no spec-defined functions
}

type Entity struct {
	Node
	publicId     string
	systemId     string
	notationName string
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

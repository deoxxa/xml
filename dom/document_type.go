package dom

type IDocumentType interface {
	INode

	// accessors
	GetName() string
	SetName(name string)
	GetEntities() INamedNodeMap
	SetEntities(entities INamedNodeMap)
	GetNotations() INamedNodeMap
	SetNotations(notations INamedNodeMap)
	GetPublicId() string
	SetPublicId(publicId string)
	GetSystemId() string
	SetSystemId(systemId string)
	GetInternalSubset() string
	SetInternalSubset(internalSubset string)
	// no spec-defined functions
}

type DocumentType struct {
	Node
	name           string
	entities       INamedNodeMap
	notations      INamedNodeMap
	publicId       string
	systemId       string
	internalSubset string
}

func (d DocumentType) GetName() string {
	return d.name
}

func (d DocumentType) SetName(name string) {
	d.name = name
}

func (d DocumentType) GetEntities() INamedNodeMap {
	return d.entities
}

func (d DocumentType) SetEntities(entities INamedNodeMap) {
	d.entities = entities
}

func (d DocumentType) GetNotations() INamedNodeMap {
	return d.notations
}

func (d DocumentType) SetNotations(notations INamedNodeMap) {
	d.notations = notations
}

func (d DocumentType) GetPublicId() string {
	return d.publicId
}

func (d DocumentType) SetPublicId(publicId string) {
	d.publicId = publicId
}

func (d DocumentType) GetSystemId() string {
	return d.systemId
}

func (d DocumentType) SetSystemId(systemId string) {
	d.systemId = systemId
}

func (d DocumentType) GetInternalSubset() string {
	return d.internalSubset
}

func (d DocumentType) SetInternalSubset(internalSubset string) {
	d.internalSubset = internalSubset
}

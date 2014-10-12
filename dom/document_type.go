package dom

type DocumentType interface {
	Node

	// accessors
	GetName() string
	SetName(name string)
	GetEntities() NamedNodeMap
	SetEntities(entities NamedNodeMap)
	GetNotations() NamedNodeMap
	SetNotations(notations NamedNodeMap)
	GetPublicId() string
	SetPublicId(publicId string)
	GetSystemId() string
	SetSystemId(systemId string)
	GetInternalSubset() string
	SetInternalSubset(internalSubset string)
	// no spec-defined functions
}

type DocumentTypeImpl struct {
	NodeImpl
	name           string
	entities       NamedNodeMap
	notations      NamedNodeMap
	publicId       string
	systemId       string
	internalSubset string
}

func (d *DocumentTypeImpl) GetName() string {
	return d.name
}

func (d *DocumentTypeImpl) SetName(name string) {
	d.name = name
}

func (d *DocumentTypeImpl) GetEntities() NamedNodeMap {
	return d.entities
}

func (d *DocumentTypeImpl) SetEntities(entities NamedNodeMap) {
	d.entities = entities
}

func (d *DocumentTypeImpl) GetNotations() NamedNodeMap {
	return d.notations
}

func (d *DocumentTypeImpl) SetNotations(notations NamedNodeMap) {
	d.notations = notations
}

func (d *DocumentTypeImpl) GetPublicId() string {
	return d.publicId
}

func (d *DocumentTypeImpl) SetPublicId(publicId string) {
	d.publicId = publicId
}

func (d *DocumentTypeImpl) GetSystemId() string {
	return d.systemId
}

func (d *DocumentTypeImpl) SetSystemId(systemId string) {
	d.systemId = systemId
}

func (d *DocumentTypeImpl) GetInternalSubset() string {
	return d.internalSubset
}

func (d *DocumentTypeImpl) SetInternalSubset(internalSubset string) {
	d.internalSubset = internalSubset
}

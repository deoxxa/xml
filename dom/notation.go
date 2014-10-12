package dom

type Notation interface {
	Node

	// accessors
	GetPublicId() string
	SetPublicId(publicId string)
	GetSystemId() string
	SetSystemId(systemId string)
	// no spec-defined functions
}

type NotationImpl struct {
	NodeImpl
	publicId string
	systemId string
}

func (n *NotationImpl) GetPublicId() string {
	return n.publicId
}

func (n *NotationImpl) SetPublicId(publicId string) {
	n.publicId = publicId
}

func (n *NotationImpl) GetSystemId() string {
	return n.systemId
}

func (n *NotationImpl) SetSystemId(systemId string) {
	n.systemId = systemId
}

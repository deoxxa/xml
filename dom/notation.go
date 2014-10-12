package dom

type INotation interface {
	INode

	// accessors
	GetPublicId() string
	SetPublicId(publicId string)
	GetSystemId() string
	SetSystemId(systemId string)
	// no spec-defined functions
}

type Notation struct {
	Node
	publicId string
	systemId string
}

func (n *Notation) GetPublicId() string {
	return n.publicId
}

func (n *Notation) SetPublicId(publicId string) {
	n.publicId = publicId
}

func (n *Notation) GetSystemId() string {
	return n.systemId
}

func (n *Notation) SetSystemId(systemId string) {
	n.systemId = systemId
}

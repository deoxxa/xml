package dom

type IDocumentFragment interface {
	INode

	// no accessors
	// no spec-defined functions
}

type DocumentFragment struct {
	Node
}

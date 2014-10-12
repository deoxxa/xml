package dom

type EntityReference interface {
	Node

	// no accessors
	// no spec-defined functions
}

type EntityReferenceImpl struct {
	NodeImpl
}

package dom

type DocumentFragment interface {
	Node

	// no accessors
	// no spec-defined functions
}

type DocumentFragmentImpl struct {
	NodeImpl
}

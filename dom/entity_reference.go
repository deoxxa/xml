package dom

type IEntityReference interface {
	INode

	// no accessors
	// no spec-defined functions
}

type EntityReference struct {
	Node
}

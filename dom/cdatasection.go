package dom

type CDATASection interface {
	Text

	// no accessors
	// no spec-defined functions
}

type CDATASectionImpl struct {
	TextImpl
}

package dom

type ICDATASection interface {
	IText

	// no accessors
	// no spec-defined functions
}

type CDATASection struct {
	Text
}

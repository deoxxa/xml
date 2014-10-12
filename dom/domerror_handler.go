package dom

type IDOMErrorHandler interface {
	// no accessors
	// spec-defined functions
	HandleError(error IDOMError) bool
}

type DOMErrorHandler struct {
}

func (d DOMErrorHandler) HandleError(error IDOMError) bool {
	panic("unimplemented") // TODO

	return false
}

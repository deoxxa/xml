package dom

type DOMErrorHandler interface {
	// no accessors
	// spec-defined functions
	HandleError(error DOMError) bool
}

type DOMErrorHandlerImpl struct {
}

func (d *DOMErrorHandlerImpl) HandleError(error DOMError) bool {
	panic("unimplemented") // TODO

	return false
}

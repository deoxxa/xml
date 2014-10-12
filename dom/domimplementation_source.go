package dom

type DOMImplementationSource interface {
	// no accessors
	// spec-defined functions
	GetDOMImplementation(features string) DOMImplementation
	GetDOMImplementationList(features string) DOMImplementationList
}

type DOMImplementationSourceImpl struct {
}

func (d DOMImplementationSourceImpl) GetDOMImplementation(features string) DOMImplementation {
	panic("unimplemented") // TODO

	return nil
}

func (d DOMImplementationSourceImpl) GetDOMImplementationList(features string) DOMImplementationList {
	panic("unimplemented") // TODO

	return nil
}

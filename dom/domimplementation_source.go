package dom

type IDOMImplementationSource interface {
	// no accessors
	// spec-defined functions
	GetDOMImplementation(features string) IDOMImplementation
	GetDOMImplementationList(features string) IDOMImplementationList
}

type DOMImplementationSource struct {
}

func (d DOMImplementationSource) GetDOMImplementation(features string) IDOMImplementation {
	panic("unimplemented") // TODO

	return nil
}

func (d DOMImplementationSource) GetDOMImplementationList(features string) IDOMImplementationList {
	panic("unimplemented") // TODO

	return nil
}

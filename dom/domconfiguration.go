package dom

type IDOMConfiguration interface {
	// accessors
	GetParameterNames() IDOMStringList
	SetParameterNames(parameterNames IDOMStringList)
	// spec-defined functions
	SetParameter(name string, value IDOMUserData)
	GetParameter(name string) IDOMUserData
	CanSetParameter(name string, value IDOMUserData) bool
}

type DOMConfiguration struct {
	parameterNames IDOMStringList
}

func (d *DOMConfiguration) GetParameterNames() IDOMStringList {
	return d.parameterNames
}

func (d *DOMConfiguration) SetParameterNames(parameterNames IDOMStringList) {
	d.parameterNames = parameterNames
}

func (d *DOMConfiguration) SetParameter(name string, value IDOMUserData) {
	panic("unimplemented") // TODO

	return
}

func (d *DOMConfiguration) GetParameter(name string) IDOMUserData {
	panic("unimplemented") // TODO

	return nil
}

func (d *DOMConfiguration) CanSetParameter(name string, value IDOMUserData) bool {
	panic("unimplemented") // TODO

	return false
}

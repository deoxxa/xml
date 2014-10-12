package dom

type DOMConfiguration interface {
	// accessors
	GetParameterNames() DOMStringList
	SetParameterNames(parameterNames DOMStringList)
	// spec-defined functions
	SetParameter(name string, value DOMUserData)
	GetParameter(name string) DOMUserData
	CanSetParameter(name string, value DOMUserData) bool
}

type DOMConfigurationImpl struct {
	parameterNames DOMStringList
}

func (d *DOMConfigurationImpl) GetParameterNames() DOMStringList {
	return d.parameterNames
}

func (d *DOMConfigurationImpl) SetParameterNames(parameterNames DOMStringList) {
	d.parameterNames = parameterNames
}

func (d *DOMConfigurationImpl) SetParameter(name string, value DOMUserData) {
	panic("unimplemented") // TODO

	return
}

func (d *DOMConfigurationImpl) GetParameter(name string) DOMUserData {
	panic("unimplemented") // TODO

	return nil
}

func (d *DOMConfigurationImpl) CanSetParameter(name string, value DOMUserData) bool {
	panic("unimplemented") // TODO

	return false
}

package dom

type Attr interface {
	Node

	// accessors
	GetName() string
	SetName(name string)
	GetSpecified() bool
	SetSpecified(specified bool)
	GetValue() string
	SetValue(value string)
	GetOwnerElement() Element
	SetOwnerElement(ownerElement Element)
	GetSchemaTypeInfo() TypeInfo
	SetSchemaTypeInfo(schemaTypeInfo TypeInfo)
	GetIsId() bool
	SetIsId(isId bool)
	// no spec-defined functions
}

type AttrImpl struct {
	NodeImpl
	name           string
	specified      bool
	value          string
	ownerElement   Element
	schemaTypeInfo TypeInfo
	isId           bool
}

func (a *AttrImpl) GetName() string {
	return a.name
}

func (a *AttrImpl) SetName(name string) {
	a.name = name
}

func (a *AttrImpl) GetSpecified() bool {
	return a.specified
}

func (a *AttrImpl) SetSpecified(specified bool) {
	a.specified = specified
}

func (a *AttrImpl) GetValue() string {
	return a.value
}

func (a *AttrImpl) SetValue(value string) {
	a.value = value
}

func (a *AttrImpl) GetOwnerElement() Element {
	return a.ownerElement
}

func (a *AttrImpl) SetOwnerElement(ownerElement Element) {
	a.ownerElement = ownerElement
}

func (a *AttrImpl) GetSchemaTypeInfo() TypeInfo {
	return a.schemaTypeInfo
}

func (a *AttrImpl) SetSchemaTypeInfo(schemaTypeInfo TypeInfo) {
	a.schemaTypeInfo = schemaTypeInfo
}

func (a *AttrImpl) GetIsId() bool {
	return a.isId
}

func (a *AttrImpl) SetIsId(isId bool) {
	a.isId = isId
}

package dom

type TypeInfo interface {
	// accessors
	GetTypeName() string
	SetTypeName(typeName string)
	GetTypeNamespace() string
	SetTypeNamespace(typeNamespace string)
	// spec-defined functions
	IsDerivedFrom(typeNamespaceArg string, typeNameArg string, derivationMethod int32) bool
}

const (
	DERIVATION_RESTRICTION = 0x00000001
	DERIVATION_EXTENSION   = 0x00000002
	DERIVATION_UNION       = 0x00000004
	DERIVATION_LIST        = 0x00000008
)

type TypeInfoImpl struct {
	typeName      string
	typeNamespace string
}

func (t *TypeInfoImpl) GetTypeName() string {
	return t.typeName
}

func (t *TypeInfoImpl) SetTypeName(typeName string) {
	t.typeName = typeName
}

func (t *TypeInfoImpl) GetTypeNamespace() string {
	return t.typeNamespace
}

func (t *TypeInfoImpl) SetTypeNamespace(typeNamespace string) {
	t.typeNamespace = typeNamespace
}

func (t *TypeInfoImpl) IsDerivedFrom(typeNamespaceArg string, typeNameArg string, derivationMethod int32) bool {
	panic("unimplemented") // TODO

	return false
}

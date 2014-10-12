package dom

type ITypeInfo interface {
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

type TypeInfo struct {
	typeName      string
	typeNamespace string
}

func (t *TypeInfo) GetTypeName() string {
	return t.typeName
}

func (t *TypeInfo) SetTypeName(typeName string) {
	t.typeName = typeName
}

func (t *TypeInfo) GetTypeNamespace() string {
	return t.typeNamespace
}

func (t *TypeInfo) SetTypeNamespace(typeNamespace string) {
	t.typeNamespace = typeNamespace
}

func (t *TypeInfo) IsDerivedFrom(typeNamespaceArg string, typeNameArg string, derivationMethod int32) bool {
	panic("unimplemented") // TODO

	return false
}

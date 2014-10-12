package main

type IDL struct {
	Interfaces []InterfaceDef
}

type InterfaceDef struct {
	Name       string
	Inherits   string
	Attributes []AttributeDef
	Functions  []FunctionDef
	Constants  []ConstantDef
}

type AttributeDef struct {
	ReadOnly bool
	Type     TypeDef
	Name     string
}

type ConstantDef struct {
	Type  TypeDef
	Name  string
	Value string
}

type FunctionDef struct {
	Type      TypeDef
	Name      string
	Arguments []ArgumentDef
}

type ArgumentDef struct {
	Type TypeDef
	Name string
}

type TypeDef struct {
	Unsigned bool
	Type     string
}

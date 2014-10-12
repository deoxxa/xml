package dom

type ICharacterData interface {
	INode

	// accessors
	GetData() string
	SetData(data string)
	GetLength() int32
	SetLength(length int32)
	// spec-defined functions
	SubstringData(offset int32, count int32) string
	AppendData(arg string)
	InsertData(offset int32, arg string)
	DeleteData(offset int32, count int32)
	ReplaceData(offset int32, count int32, arg string)
}

type CharacterData struct {
	Node
	data   string
	length int32
}

func (c *CharacterData) GetData() string {
	return c.data
}

func (c *CharacterData) SetData(data string) {
	c.data = data
}

func (c *CharacterData) GetLength() int32 {
	return c.length
}

func (c *CharacterData) SetLength(length int32) {
	c.length = length
}

func (c *CharacterData) SubstringData(offset int32, count int32) string {
	panic("unimplemented") // TODO

	return ""
}

func (c *CharacterData) AppendData(arg string) {
	panic("unimplemented") // TODO

	return
}

func (c *CharacterData) InsertData(offset int32, arg string) {
	panic("unimplemented") // TODO

	return
}

func (c *CharacterData) DeleteData(offset int32, count int32) {
	panic("unimplemented") // TODO

	return
}

func (c *CharacterData) ReplaceData(offset int32, count int32, arg string) {
	panic("unimplemented") // TODO

	return
}

package dom

type CharacterData interface {
	Node

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

type CharacterDataImpl struct {
	NodeImpl
	data   string
	length int32
}

func (c *CharacterDataImpl) GetData() string {
	return c.data
}

func (c *CharacterDataImpl) SetData(data string) {
	c.data = data
}

func (c *CharacterDataImpl) GetLength() int32 {
	return c.length
}

func (c *CharacterDataImpl) SetLength(length int32) {
	c.length = length
}

func (c *CharacterDataImpl) SubstringData(offset int32, count int32) string {
	panic("unimplemented") // TODO

	return ""
}

func (c *CharacterDataImpl) AppendData(arg string) {
	panic("unimplemented") // TODO

	return
}

func (c *CharacterDataImpl) InsertData(offset int32, arg string) {
	panic("unimplemented") // TODO

	return
}

func (c *CharacterDataImpl) DeleteData(offset int32, count int32) {
	panic("unimplemented") // TODO

	return
}

func (c *CharacterDataImpl) ReplaceData(offset int32, count int32, arg string) {
	panic("unimplemented") // TODO

	return
}

package dom

type UserDataHandler interface {
	// no accessors
	// spec-defined functions
	Handle(operation int16, key string, data DOMUserData, src Node, dst Node)
}

const (
	NODE_CLONED   = 1
	NODE_IMPORTED = 2
	NODE_DELETED  = 3
	NODE_RENAMED  = 4
	NODE_ADOPTED  = 5
)

type UserDataHandlerImpl struct {
}

func (u *UserDataHandlerImpl) Handle(operation int16, key string, data DOMUserData, src Node, dst Node) {
	panic("unimplemented") // TODO

	return
}

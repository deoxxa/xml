package dom

type IUserDataHandler interface {
	// no accessors
	// spec-defined functions
	Handle(operation int16, key string, data IDOMUserData, src INode, dst INode)
}

const (
	NODE_CLONED   = 1
	NODE_IMPORTED = 2
	NODE_DELETED  = 3
	NODE_RENAMED  = 4
	NODE_ADOPTED  = 5
)

type UserDataHandler struct {
}

func (u *UserDataHandler) Handle(operation int16, key string, data IDOMUserData, src INode, dst INode) {
	panic("unimplemented") // TODO

	return
}

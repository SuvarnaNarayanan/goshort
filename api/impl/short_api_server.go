package impl

type GoShortApiServer struct {
	*ShortApiImpl
}

func NewGoShortApiServer() *GoShortApiServer {
	return &GoShortApiServer{
		&ShortApiImpl{},
	}
}

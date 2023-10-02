package v1

type homeHandler struct {
}

func NewHomeHandler() *homeHandler {
	return &homeHandler{}
}

func (*homeHandler) homePage() {

}

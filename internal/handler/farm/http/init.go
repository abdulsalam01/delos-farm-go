package http

func New(usecase Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

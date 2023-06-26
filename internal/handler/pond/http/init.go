package http

func New(
	farmUsecase FarmUsecase,
	pondUsecase PondUsecase,
) *Handler {
	return &Handler{
		farmUsecase: farmUsecase,
		pondUsecase: pondUsecase,
	}
}

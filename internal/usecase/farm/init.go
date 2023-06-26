package farm

func New(
	farmRepo farmRepoResource,
	pondRepo pondRepoResouce,
) *Usecase {
	return &Usecase{
		farmRepo: farmRepo,
		pondRepo: pondRepo,
	}
}

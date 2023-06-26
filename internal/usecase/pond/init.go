package pond

func New(
	farmRepo farmRepoResource,
	pondRepo pondRepoResource,
) *Usecase {
	return &Usecase{
		farmRepo: farmRepo,
		pondRepo: pondRepo,
	}
}

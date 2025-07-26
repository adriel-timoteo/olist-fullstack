package repository

type CustomerRepoItf interface {
}

type CustomerRepoImpl struct {
}

func NewCustomerRepo() CustomerRepoImpl {
	return CustomerRepoImpl{}
}

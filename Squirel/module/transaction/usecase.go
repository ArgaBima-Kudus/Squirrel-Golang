package transaction

import (
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
)

type UseCase struct {
	Repo RepositoryInterface
}

type UseCaseInterface interface {
	GetAllTransaction(req *dto.Request) ([]entities.Transaction, error, int64)
}

func NewUseCase(repo RepositoryInterface) UseCaseInterface {
	return UseCase{
		Repo: repo,
	}
}

func (uc UseCase) GetAllTransaction(req *dto.Request) ([]entities.Transaction, error, int64) {

	return uc.Repo.GetAllTransaction(req)
}

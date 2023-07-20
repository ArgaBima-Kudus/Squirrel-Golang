package transaction

import (
	"golang/module/transaction/dto"
)

type Controller struct {
	UseCase UseCaseInterface
}

type ControllerInterface interface {
	GetAllTransactionByRequest(req *dto.Request) (*dto.GetAllResponseDataTransaction, error, int64)
}

func NewController(useCase UseCaseInterface) ControllerInterface {
	return Controller{
		UseCase: useCase,
	}
}

func (c Controller) GetAllTransactionByRequest(req *dto.Request) (*dto.GetAllResponseDataTransaction, error, int64) {
	transactions, err, totaldata := c.UseCase.GetAllTransaction(req)
	if err != nil {
		return nil, err, 0
	}

	res := &dto.GetAllResponseDataTransaction{
		Code:      200,
		Message:   "Data Berhasil Diambil",
		Error:     "Success",
		TotalData: int(totaldata),
		TotalDataFiltered: len(transactions),
	}

	for _, transaction := range transactions {
		res.Data = append(res.Data, dto.TransactionItemResponse{
			Id:         transaction.Id,
			OdaNumber:  transaction.OdaNumber,
			BillAmount: transaction.BillAmount,
			Status:     transaction.Status,
			CreatedAt:  transaction.CreatedAt,
		})
	}
	return res, nil, totaldata
}

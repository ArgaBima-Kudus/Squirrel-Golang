package exportcsv

import (
	"errors"
	"strconv"
	"strings"
)

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *useCase {
	return &useCase{
		repo: repo,
	}
}

type Usecase interface {
	ExportCSV(req *ExportCSVRequest) ([][]string, error)
	Get() ([]Transaction, error)
}

func (u useCase) Get() ([]Transaction, error) {
	data, err := u.repo.GetAllTransaction()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u useCase) ExportCSV(req *ExportCSVRequest) ([][]string, error) {

	var resultGetTransaction []Transaction
	var transactionStringData [][]string
	var err error
	switch {
	case req.Status == "" && req.StartDate == "" && req.EndDate == "": //without filter
		resultGetTransaction, err = u.repo.GetAllTransaction()
	case req.Status == "": // rage date filter
		resultGetTransaction, err = u.repo.GetAllTransactionByRangeDateFilter(req)
	case req.StartDate == "" && req.EndDate == "": //only status filter
		resultGetTransaction, err = u.repo.GetTransactionByStatusFilter(req)
	default: //complate filter(status, rage date)
		resultGetTransaction, err = u.repo.GetTransactionByStatusAndRangeDateFilter(req)
	}
	if len(resultGetTransaction) == 0 {
		return nil, errors.New("Not Found")
	}
	transactionStringData, err = TransactionStringConverter(resultGetTransaction)
	if err != nil {
		return nil, err
	}
	return transactionStringData, nil
}

func TransactionStringConverter(transactions []Transaction) ([][]string, error) {
	//set name field transaction in first record
	stringData := [][]string{[]string{"id", "oda_number", "status", "price", "created_at"}}
	//converting some data from transaction to string
	for _, transaction := range transactions {
		record := []string{strconv.Itoa(transaction.Id), transaction.Oda_number, transaction.Status, strconv.FormatFloat(transaction.Bill_amount, 'f', 6, 64), strings.Trim(transaction.Created_at, "T00:00:00Z")}
		stringData = append(stringData, record)
	}
	//
	return stringData, nil
}

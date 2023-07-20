package exportcsv

import (
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetAllTransaction() ([]Transaction, error)
	GetTransactionByStatusFilter(req *ExportCSVRequest) ([]Transaction, error)
	GetAllTransactionByRangeDateFilter(req *ExportCSVRequest) ([]Transaction, error)
	GetTransactionByStatusAndRangeDateFilter(req *ExportCSVRequest) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetAllTransaction() ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		fmt.Println(err)
		return nil, err

	}
	//	fmt.Println(transactions)
	return transactions, nil
}
func (r repository) GetTransactionByStatusFilter(req *ExportCSVRequest) ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Where("status = ?", req.Status).Find(&transactions).Error; err != nil {
		return nil, err

	}

	return transactions, nil
}
func (r repository) GetAllTransactionByRangeDateFilter(req *ExportCSVRequest) ([]Transaction, error) {
	var transactions []Transaction
	if err := r.db.Where("created_at BETWEEN  ? AND  ?", req.StartDate, req.EndDate).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r repository) GetTransactionByStatusAndRangeDateFilter(req *ExportCSVRequest) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("status =? AND(created_at  BETWEEN ? AND ?)", req.Status, req.StartDate, req.EndDate).Find(&transactions).Error
	if err != nil {
		return nil, err

	}

	return transactions, nil
}

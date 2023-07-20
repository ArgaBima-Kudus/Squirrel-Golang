package transaction

import (
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
	"golang/module/transaction/helper"

	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	GetAllTransaction(req *dto.Request) ([]entities.Transaction, error, int64)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return Repository{db}
}

func (r Repository) GetAllTransaction(req *dto.Request) ([]entities.Transaction, error, int64) {

	var count int64

	query := squirrel.Select("*").From("transaction")
	r.DB.Raw("SELECT COUNT(*) FROM transaction").Count(&count)

	if req.Page != 0 && req.Size != 0 {
		offset := uint64((req.Page - 1) * req.Size)
		limit := uint64(req.Size)
		query = query.Offset(offset).Limit(limit)

		switch {
		case req.Status != "" && req.StartDate == "" && req.EndDate == "":
			query = query.Where(squirrel.Eq{"status": req.Status})
		case req.Status == "" && req.StartDate != "" && req.EndDate != "":
			start := helper.FormatDate(req.StartDate)
			end := helper.FormatDate(req.EndDate)
			query = query.Where(squirrel.Expr("? BETWEEN created_at AND ?", start, end))
		case req.Status != "" && req.StartDate != "" && req.EndDate != "":
			start := helper.FormatDate(req.StartDate)
			end := helper.FormatDate(req.EndDate)
			query = query.Where(squirrel.Eq{"status": req.Status})
			query = query.Where(squirrel.Expr("? BETWEEN created_at AND ?", start, end))
		}
	} else {
		switch {
		case req.Status != "" && req.StartDate == "" && req.EndDate == "":
			query = query.Where(squirrel.Eq{"status": req.Status})
		case req.Status == "" && req.StartDate != "" && req.EndDate != "":
			start := helper.FormatDate(req.StartDate)
			end := helper.FormatDate(req.EndDate)
			query = query.Where(squirrel.Expr("? BETWEEN created_at AND ?", start, end))
		case req.Status != "" && req.StartDate != "" && req.EndDate != "":
			start := helper.FormatDate(req.StartDate)
			end := helper.FormatDate(req.EndDate)
			query = query.Where(squirrel.Eq{"status": req.Status})
			query = query.Where(squirrel.Expr("? BETWEEN created_at AND ?", start, end))
		}
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Question).ToSql()
	if err != nil {
		return nil, err, 0
	}

	var transactions []entities.Transaction
	err = r.DB.Raw(sql, args...).Scan(&transactions).Error
	return transactions, err, count
}

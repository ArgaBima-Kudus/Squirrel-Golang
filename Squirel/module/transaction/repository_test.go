package transaction

import (
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
	"golang/module/transaction/helper"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_GetAllTransaction(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	queryCount := "SELECT COUNT(*) FROM transaction"
	queryGetAll := "SELECT * FROM transaction"
	
	
	mockQuery.ExpectQuery(queryCount).WillReturnRows(sqlmock.NewRows([]string{"total_data"}).AddRow(1))
	mockQuery.ExpectQuery(queryGetAll).WillReturnRows(sqlmock.NewRows(
		[]string{"id", "oda_number", "bill_amount", "status", "created_at"}).
		AddRow(1, "123456789", 10000, "SUCCESS", "2023-04-08 00:00:00"),
	)

	type args struct {
		req *dto.Request
	}
	tests := []struct {
		name  string
		r     Repository
		args  args
		want  []entities.Transaction
		want1 error
		want2 int64
	}{
		{
			name: "Test Case 1",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				req: &dto.Request{
					Status:    "",
					StartDate: "",
					EndDate:   "",
					Page:      0,
					Size:      0,
				},
			},
			want: []entities.Transaction{
				{
					Id:         1,
					OdaNumber:  "123456789",
					BillAmount: 10000,
					Status:     "SUCCESS",
					CreatedAt:  "2023-04-08 00:00:00",
				},
			},
			want1: nil,
			want2: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.r.GetAllTransaction(tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransaction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetAllTransaction() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetAllTransaction() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

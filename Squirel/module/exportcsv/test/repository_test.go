package test

import (
	"golang/module/exportcsv"
	"golang/module/exportcsv/test/helper"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repository_GetAllTransaction(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")
	mockQuery.ExpectQuery("SELECT * FROM `transaction`").WillReturnRows(rows)

	tests := []struct {
		name    string
		r       exportcsv.Repository
		want    []exportcsv.Transaction
		wantErr bool
	}{
		{name: "succes",
			r: exportcsv.NewRepository(mockDb),
			want: []exportcsv.Transaction{
				{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllTransaction()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAllTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetTransactionByStatusFilter(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")

	mockQuery.ExpectQuery("SELECT * FROM `transaction` WHERE status = ?").WillReturnRows(rows)

	type args struct {
		req *exportcsv.ExportCSVRequest
	}
	tests := []struct {
		name    string
		r       exportcsv.Repository
		args    args
		want    []exportcsv.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			r:    exportcsv.NewRepository(mockDb),
			args: args{req: &exportcsv.ExportCSVRequest{Status: "SUCCESS", StartDate: "", EndDate: ""}},
			want: []exportcsv.Transaction{
				{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetTransactionByStatusFilter(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetTransactionByStatusFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetTransactionByStatusFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetAllTransactionByRangeDateFilter(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")

	mockQuery.ExpectQuery("SELECT * FROM `transaction` WHERE created_at BETWEEN ? AND ?").WillReturnRows(rows)

	type args struct {
		req *exportcsv.ExportCSVRequest
	}
	tests := []struct {
		name    string
		r       exportcsv.Repository
		args    args
		want    []exportcsv.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			r:    exportcsv.NewRepository(mockDb),
			args: args{req: &exportcsv.ExportCSVRequest{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}},
			want: []exportcsv.Transaction{
				{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllTransactionByRangeDateFilter(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllTransactionByRangeDateFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAllTransactionByRangeDateFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetTransactionByStatusAndRangeDateFilter(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")

	mockQuery.ExpectQuery("SELECT * FROM `transaction` WHERE status =? AND(created_at BETWEEN ? AND ?)").WillReturnRows(rows)

	type args struct {
		req *exportcsv.ExportCSVRequest
	}
	tests := []struct {
		name    string
		r       exportcsv.Repository
		args    args
		want    []exportcsv.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			r:    exportcsv.NewRepository(mockDb),
			args: args{req: &exportcsv.ExportCSVRequest{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}},
			want: []exportcsv.Transaction{
				{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
				{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetTransactionByStatusAndRangeDateFilter(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetTransactionByStatusAndRangeDateFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetTransactionByStatusAndRangeDateFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

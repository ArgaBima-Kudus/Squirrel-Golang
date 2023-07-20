package transaction

import (
	"golang/module/transaction/mocks"
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
	"reflect"
	"testing"
)

func TestUseCase_GetAllTransaction(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
	}

	type args struct {
		req *dto.Request
	}
	tests := []struct {
		name    string
		field   fields
		args    args
		want    []entities.Transaction
		want1   int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			field: fields{
				Repo: mocks.NewRepositoryInterface(t),
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
			want1:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			U := UseCase{
				Repo: tt.field.Repo,
			}
			// Menetapkan tipe kembalian untuk mock Repo
			tt.field.Repo.(*mocks.RepositoryInterface).On("GetAllTransaction", tt.args.req).Return(tt.want, nil, tt.want1)
			got, err, got1 := U.GetAllTransaction(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetAllTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.GetAllTransaction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UseCase.GetAllTransaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

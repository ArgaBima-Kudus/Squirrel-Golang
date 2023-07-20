package transaction

import (
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
	"golang/module/transaction/mocks"
	"reflect"
	"testing"
)

func TestController_GetAllTransactionByRequest(t *testing.T) {

	mockController := mocks.NewUseCaseInterface(t)
	mockController.EXPECT().GetAllTransaction(
		&dto.Request{
			Status:    "",
			StartDate: "",
			EndDate:   "",
			Page:      0,
			Size:      0,
		},
	).Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "123456789",
			BillAmount: 10000,
			Status:     "SUCCESS",
			CreatedAt:  "2023-04-08 00:00:00",
		},
	}, nil, 1)

	type fields struct {
		UseCase UseCaseInterface
	}

	type args struct {
		req *dto.Request
	}
	tests := []struct {
		name    string
		fields fields
		args    args
		want    *dto.GetAllResponseDataTransaction
		want1   int64
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				UseCase: mockController,
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
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				TotalDataFiltered: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "123456789",
						BillAmount: 10000,
						Status:     "SUCCESS",
						CreatedAt:  "2023-04-08 00:00:00",
					},
				},
			},
			want1:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.fields.UseCase,
			}
			got, err, got1 := c.GetAllTransactionByRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.GetAllTransactionByRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.GetAllTransactionByRequest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Controller.GetAllTransactionByRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

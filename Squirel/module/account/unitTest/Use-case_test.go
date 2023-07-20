package unitTest

import (
	"golang/module/account"
	"golang/module/account/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_GetDataUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().GetDataUser().Return([]account.Account{{
		Name:  "Vincen",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}}, nil).Times(1)
	tests := []struct {
		name    string
		u       account.UseCase
		want    []account.Account
		wantErr bool
	}{
		{
			name: "Success Get Data User",
			u: account.UseCase{
				Repo: mockery,
			},
			want: []account.Account{{
				Name:  "Vincen",
				Email: "vincen@gmail.com",
				Role:  "admin",
				Phone: "123",
			},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetDataUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.GetDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetDataUserById(t *testing.T) {
	type args struct {
		id string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().GetDataUserById("1").Return(account.Account{
		Name:  "Vincen",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}, nil).Times(1)

	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "Success Get Data User By Id",
			u:    account.UseCase{Repo: mockery},
			args: args{id: "1"},
			want: account.Account{
				Name:  "Vincen",
				Email: "vincen@gmail.com",
				Role:  "admin",
				Phone: "123",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetDataUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetDataUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.GetDataUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseCase_EditDataUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().EditDataUser("1", &account.Account{
		Name:  "Vincencius Maxwell",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}).Return(account.Account{
		Name:  "Vincencius Maxwell",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}, nil).Times(1)

	type args struct {
		id  string
		req *account.Account
	}
	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "Success Edit Data User",
			u:    account.UseCase{Repo: mockery},
			args: args{
				id: "1",
				req: &account.Account{
					Name:  "Vincencius Maxwell",
					Email: "vincen@gmail.com",
					Role:  "admin",
					Phone: "123",
				},
			},
			want: account.Account{
				Name:  "Vincencius Maxwell",
				Email: "vincen@gmail.com",
				Role:  "admin",
				Phone: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.EditDataUser(tt.args.id, tt.args.req)

			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.EditDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.EditDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseCase_CreateAccount(t *testing.T) {
	type args struct {
		req *account.Account
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().CreateAccount(&account.Account{
		Id:       1,
		Name:     "Vincencius Maxwell",
		Email:    "vincen@gmail.com",
		Password: "123456",
		Role:     "admin",
		Phone:    "123",
	}).Return(account.Account{
		Id:       1,
		Name:     "Vincencius Maxwell",
		Email:    "vincen@gmail.com",
		Password: "123456",
		Role:     "admin",
		Phone:    "123",
	}, nil).Times(1)

	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "Success Create Account",
			u: account.UseCase{
				Repo: mockery,
			},
			args: args{
				req: &account.Account{
					Id:       1,
					Name:     "Vincencius Maxwell",
					Email:    "vincen@gmail.com",
					Password: "123456",
					Role:     "admin",
					Phone:    "123",
				},
			},
			want: account.Account{
				Id:       1,
				Name:     "Vincencius Maxwell",
				Email:    "vincen@gmail.com",
				Password: "123456",
				Role:     "admin",
				Phone:    "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.CreateAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseCase_DeleteDataUser(t *testing.T) {
	type args struct {
		id string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().DeleteDataUser("1").Return(account.Account{}, nil).Times(1)

	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "Success Delete Data User",
			u: account.UseCase{
				Repo: mockery,
			},
			args: args{
				id: "1",
			},
			want:    account.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.DeleteDataUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.DeleteDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.DeleteDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseCase_Login(t *testing.T) {
	type args struct {
		req *account.Account
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().Login(&account.Account{
		Email:    "vincen@gmail.com",
		Password: "123456",
	},
	).Return("token", account.Account{}, nil).Times(1)

	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    string
		want1   account.Account
		wantErr bool
	}{
		{
			name: "Success Login",
			u: account.UseCase{
				Repo: mockery,
			},
			args: args{
				&account.Account{
					Email:    "vincen@gmail.com",
					Password: "123456",
				},
			},
			want:  "token",
			want1: account.Account{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.u.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.Login() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UseCase.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func TestUseCase_SendEmail(t *testing.T) {
	type args struct {
		email string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().SendEmail("maxwelvincen@gmail.com").Return(account.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "success",
			u: account.UseCase{
				Repo: mockery,
			},
			args: args{
				email: "maxwelvincen@gmail.com",
			},
			want:    account.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.SendEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.SendEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUseCase_CompareVerificationCode(t *testing.T) {
	type args struct {
		verificationCode *account.VerificationCodeRequest
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().CompareVerificationCode(&account.VerificationCodeRequest{
		Email: "maxwelvincen@gmail.com",
		Code:  "1234",
	}).Return(account.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "Success",
			u: account.UseCase{
				Repo: mockery,
			},
			args: args{
				&account.VerificationCodeRequest{
					Email: "maxwelvincen@gmail.com",
					Code:  "1234",
				},
			},
			want:    account.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.CompareVerificationCode(tt.args.verificationCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CompareVerificationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.CompareVerificationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_EditPassword(t *testing.T) {
	type args struct {
		id  string
		req *account.Account
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().EditPassword("1", &account.Account{
		Password: "123456",
	}).Return(account.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       account.UseCase
		args    args
		want    account.Account
		wantErr bool
	}{
		{
			name: "success",
			u: account.UseCase{
				Repo: mockery,
			},
			args: args{
				"1", &account.Account{
					Password: "123456",
				},
			},
			want:    account.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.EditPassword(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.EditPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.EditPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

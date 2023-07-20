package account

import (
	"net/smtp"

	"golang.org/x/crypto/bcrypt"
)

type ControllerInterface interface {
	GetDataUser() (*ReadResponse, error)
	GetDataUserById(id string) (*ReadResponse, error)
	CreateAccount(req *CreateRequest) (*CreateResponse, error)
	EditDataUser(id string, req *EditDataUserRequest) (*CreateResponse, error)
	DeleteDataUser(id string) (*CreateResponse, error)
	Login(req *LoginResponseRequest) (string, *LoginResponse, error)
	SendEmail(email string) (*CreateResponse, error)
	SendEmailRegister(email string) (*CreateResponse, error)
	CompareVerificationCode(verificationCode *VerificationCodeRequest) (*CreateResponse, error)
	EditPassword(id string, code string, req *EditDataUserRequest) (*CreateResponse, error)
}

type Controller struct {
	UseCase UseCaseInterface
}

func NewController(useCase UseCaseInterface) ControllerInterface {
	return Controller{
		UseCase: useCase,
	}
}

type CreateResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Status  string                `json:"status"`
	Data    []AccountItemResponse `json:"data"`
}
type ReadResponse struct {
	Message string                `json:"message"`
	Code    int                   `json:"code"`
	Status  string                `json:"status"`
	Data    []AccountItemResponse `json:"data"`
}
type AccountItemResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
type LoginResponse struct {
	Message string                   `json:"message"`
	Code    int                      `json:"code"`
	Status  string                   `json:"status"`
	Data    []LoginResponseWithToken `json:"data"`
}
type LoginResponseWithToken struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (c Controller) GetDataUser() (*ReadResponse, error) {
	account, err := c.UseCase.GetDataUser()
	if err != nil {
		res := &ReadResponse{
			Status:  "Error",
			Message: "Internal Server Error",
			Code:    400,
			Data:    nil,
		}
		return res, nil
	}

	res := &ReadResponse{
		Status:  "OK",
		Message: "Success Get Data",
		Code:    200,
	}

	for _, account := range account {
		item := AccountItemResponse{
			Id:    account.Id,
			Name:  account.Name,
			Role:  account.Role,
			Email: account.Email,
			Phone: account.Phone,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) GetDataUserById(id string) (*ReadResponse, error) {
	account, err := c.UseCase.GetDataUserById(id)
	if err != nil {
		return nil, err
	}
	if account.Id == 0 {
		res := &ReadResponse{
			Code:    204,
			Status:  "OK",
			Message: "Data With ID: " + id + " Empty",
			Data:    nil,
		}
		return res, nil
	}
	res := &ReadResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Get Data",
		Data: []AccountItemResponse{{
			Id:    account.Id,
			Name:  account.Name,
			Role:  account.Role,
			Email: account.Email,
			Phone: account.Phone,
		},
		},
	}
	return res, nil
}
func (c Controller) CreateAccount(req *CreateRequest) (*CreateResponse, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := Account{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	_, err := c.UseCase.CreateAccount(&request)
	if err != nil {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Duplicate account name",
			Data:    nil,
		}
		return res, err
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Save Data",
		Data: []AccountItemResponse{{
			Id:    request.Id,
			Name:  request.Name,
			Role:  request.Role,
			Email: request.Email,
			Phone: request.Phone,
		},
		},
	}
	return res, nil

}
func (c Controller) EditDataUser(id string, req *EditDataUserRequest) (*CreateResponse, error) {

	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := Account{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	_, err := c.UseCase.EditDataUser(id, &request)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Edit Data",
		Data: []AccountItemResponse{{
			Id:    request.Id,
			Name:  request.Name,
			Role:  request.Role,
			Email: request.Email,
			Phone: request.Phone,
		},
		},
	}
	return res, nil
}
func (c Controller) DeleteDataUser(id string) (*CreateResponse, error) {
	_, err := c.UseCase.DeleteDataUser(id)
	if err != nil {
		return nil, err
	}
	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Delete Data",
	}
	return res, nil
}
func (c Controller) Login(req *LoginResponseRequest) (string, *LoginResponse, error) {
	request := Account{
		Email:    req.Email,
		Password: req.Password,
	}
	token, data, err := c.UseCase.Login(&request)
	if err != nil {
		res := &LoginResponse{
			Code:    401,
			Status:  "Error",
			Message: "Email atau Password Salah",
			Data:    nil,
		}
		return "", res, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password)); err != nil {
		res := &LoginResponse{
			Code:    401,
			Status:  "Error",
			Message: "Email atau Password Salah",
			Data:    nil,
		}
		return "", res, err
	}

	res := &LoginResponse{
		Code:    200,
		Status:  "OK",
		Message: "Login Success",
		Data: []LoginResponseWithToken{{
			Id:    data.Id,
			Name:  data.Name,
			Email: data.Email,
			Role:  data.Role,
			Token: token,
		},
		},
	}
	return token, res, nil
}
func (c Controller) SendEmail(email string) (*CreateResponse, error) {
	data, _ := c.UseCase.SendEmail(email)
	if data.Id == 0 {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Email not found",
			Data:    nil,
		}
		return res, nil
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "briceria123@gmail.com"
	password := "fjuyqpqqnrkzaatr"
	rand := GenerateVerificationCode(email)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	subject := "Subject: [Verification Code] Forgot Password Account App Dashboard DDB Ceria\n"

	mainMessage1 := "<body marginheight='0' topmargin='0' marginwidth='0' style='margin: 0px; background-color: #f2f3f8;' leftmargin='0'><table cellspacing='0' border='0' cellpadding='0' width='100%' bgcolor='#f2f3f8'style='@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;'><tr><td><table style='background-color: #f2f3f8; max-width:670px;  margin:0 auto;' width='100%' border='0'align='center' cellpadding='0' cellspacing='0'><tr><td style='height:80px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><a href='https://bri.co.id/web/ceria' title='logo' target='_blank'><img width='60' src='https://play-lh.googleusercontent.com/tpsB_EJ4_p3Ljh7LwhNWg6ysAH8GoDzDIcZwIWTP9SX1HsVjPflGP_iUK4IWGZOulDk=w480-h960-rw' title='logo' alt='logo'></a></td></tr><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td><table width='95%' border='0' align='center' cellpadding='0' cellspacing='0'style='max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);'><tr><td style='height:40px;'>&nbsp;</td></tr><tr><td style='padding:0 35px;'><h1 style='color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;'>You have requested to reset your password</h1><spanstyle='display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;'></span><p style='color:#455056; font-size:15px;line-height:24px; margin:0;'>We cannot simply send you your old password. Please copy your verification code. To reset your password</p><a href='javascript:void(0);'style='background:#20e277;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;'>"
	mainMessage2 := "</a></td></tr><tr><td style='height:40px;'>&nbsp;</td></tr></table></td><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><p style='font-size:14px; color:rgba(69, 80, 86, 0.7411764705882353); line-height:18px; margin:0 0 0;'>&copy; <strong>https://bri.co.id/web/ceria</strong></p></td></tr><tr><td style='height:80px;'>&nbsp;</td></tr></table></td></tr></table></body>"

	body := mainMessage1 + rand + mainMessage2
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		res := &CreateResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed Send Email",
			Data:    nil,
		}
		return res, nil
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Send Email",
		Data: []AccountItemResponse{{
			Id:    data.Id,
			Email: data.Email,
			Name:  data.Name,
			Role:  data.Role,
			Phone: data.Phone,
		},
		},
	}
	return res, nil
}
func (c Controller) SendEmailRegister(email string) (*CreateResponse, error) {
	data, _ := c.UseCase.SendEmail(email)
	if data.Id != 0 {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Email Already Taken",
			Data:    nil,
		}
		return res, nil
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "briceria123@gmail.com"
	password := "fjuyqpqqnrkzaatr"
	rand := GenerateVerificationCode(email)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	to := []string{email}

	subject := "Subject: [Verification Code] Registrasion Account App Dashboard DDB Ceria\n"

	mainMessage1 := "<body marginheight='0' topmargin='0' marginwidth='0' style='margin: 0px; background-color: #f2f3f8;' leftmargin='0'><table cellspacing='0' border='0' cellpadding='0' width='100%' bgcolor='#f2f3f8'style='@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;'><tr><td><table style='background-color: #f2f3f8; max-width:670px;  margin:0 auto;' width='100%' border='0'align='center' cellpadding='0' cellspacing='0'><tr><td style='height:80px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><a href='https://bri.co.id/web/ceria' title='logo' target='_blank'><img width='60' src='https://play-lh.googleusercontent.com/tpsB_EJ4_p3Ljh7LwhNWg6ysAH8GoDzDIcZwIWTP9SX1HsVjPflGP_iUK4IWGZOulDk=w480-h960-rw' title='logo' alt='logo'></a></td></tr><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td><table width='95%' border='0' align='center' cellpadding='0' cellspacing='0'style='max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);'><tr><td style='height:40px;'>&nbsp;</td></tr><tr><td style='padding:0 35px;'><h1 style='color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;'>You have requested to registration your account</h1><spanstyle='display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;'></span><p style='color:#455056; font-size:15px;line-height:24px; margin:0;'>Please copy your verification code. To registration your password</p><a href='javascript:void(0);'style='background:#20e277;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;'>"
	mainMessage2 := "</a></td></tr><tr><td style='height:40px;'>&nbsp;</td></tr></table></td><tr><td style='height:20px;'>&nbsp;</td></tr><tr><td style='text-align:center;'><p style='font-size:14px; color:rgba(69, 80, 86, 0.7411764705882353); line-height:18px; margin:0 0 0;'>&copy; <strong>https://bri.co.id/web/ceria</strong></p></td></tr><tr><td style='height:80px;'>&nbsp;</td></tr></table></td></tr></table></body>"

	body := mainMessage1 + rand + mainMessage2
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		res := &CreateResponse{
			Code:    500,
			Status:  "Error",
			Message: "Failed Send Email",
			Data:    nil,
		}
		return res, nil
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Send Email",
		Data:    nil,
	}
	return res, nil
}
func (c Controller) CompareVerificationCode(verificationCode *VerificationCodeRequest) (*CreateResponse, error) {
	data, _ := c.UseCase.CompareVerificationCode(verificationCode)

	CodeFromMap := VerificationCodes[verificationCode.Email]
	CodeFromRequest := verificationCode.Code
	if CodeFromMap != CodeFromRequest {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid Verification Code",
			Data:    nil,
		}
		return res, nil
	}
	if data.Id == 0 {
		res := &CreateResponse{
			Code:    200,
			Status:  "OK",
			Message: "Success Verification Code",
			Data:    nil,
		}
		return res, nil
	}
	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Verification Code",
		Data: []AccountItemResponse{{
			Id:    data.Id,
			Email: data.Email,
			Name:  data.Name,
			Role:  data.Role,
			Phone: data.Phone,
		},
		},
	}
	return res, nil

}
func (c Controller) EditPassword(id string, code string, req *EditDataUserRequest) (*CreateResponse, error) {
	result, err := c.UseCase.GetDataUserById(id)
	if err != nil {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Account Not Found",
			Data:    nil,
		}
		return res, nil
	}
	CodeFromMap := VerificationCodes[result.Email]
	CodeFromRequest := code

	if CodeFromMap != CodeFromRequest {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Invalid Verification Code",
			Data:    nil,
		}
		return res, nil
	}
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte((req.Password)), bcrypt.DefaultCost)
	req.Password = string(HashPassword)

	request := Account{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: req.Password,
	}
	data, _ := c.UseCase.EditPassword(id, &request)
	if data.Id == 0 {
		res := &CreateResponse{
			Code:    400,
			Status:  "Error",
			Message: "Failed to update account",
			Data:    nil,
		}
		return res, nil
	}

	res := &CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success Update Password",
		Data: []AccountItemResponse{{
			Id:    data.Id,
			Email: data.Email,
			Name:  data.Name,
			Role:  data.Role,
			Phone: data.Phone,
		},
		},
	}
	return res, nil
}

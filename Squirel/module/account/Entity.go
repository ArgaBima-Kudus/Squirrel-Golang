package account

type Account struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type VerificationCodeRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

var VerificationCodes = make(map[string]string)

package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandlerInterface interface {
	GetDataUser(c *gin.Context)
	GetDataUserById(c *gin.Context)
	CreateAccount(c *gin.Context)
	EditDataUser(c *gin.Context)
	DeleteDataUser(c *gin.Context)
	Login(c *gin.Context)
	SendEmail(c *gin.Context)
	SendEmailRegister(c *gin.Context)
	CompareVerificationCode(c *gin.Context)
	EditPassword(c *gin.Context)
}
type RequestHandler struct {
	Ctrl ControllerInterface
}

func NewRequestHandler(ctrl ControllerInterface) RequestHandlerInterface {
	return &RequestHandler{
		Ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) RequestHandlerInterface {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

type MessageResponse struct {
	Message string `json:"message"`
}
type CreateRequest struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type LoginResponseRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
type EditDataUserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h RequestHandler) GetDataUser(c *gin.Context) {
	res, err := h.Ctrl.GetDataUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h RequestHandler) GetDataUserById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Ctrl.GetDataUserById(id)
	if err != nil {
		c.JSON(500, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}
func (h RequestHandler) CreateAccount(c *gin.Context) {
	var account CreateRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	res, _ := h.Ctrl.CreateAccount(&account)
	c.JSON(200, res)

}

func (h RequestHandler) EditDataUser(c *gin.Context) {
	var account EditDataUserRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	res, err := h.Ctrl.EditDataUser(id, &account)
	if err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, res)
}
func (h RequestHandler) DeleteDataUser(c *gin.Context) {
	id := c.Param("id")
	res, err := h.Ctrl.DeleteDataUser(id)
	if err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
	}
	c.JSON(200, res)
}
func (h RequestHandler) Login(c *gin.Context) {
	var account LoginResponseRequest

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, ErrorResponse{Error: err.Error()})
		return
	}
	token, res, err := h.Ctrl.Login(&account)
	if err != nil {
		c.JSON(401, res)
		return
	}
	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, res)
}

func (h RequestHandler) SendEmail(c *gin.Context) {
	email := c.Param("email")
	res, err := h.Ctrl.SendEmail(email)
	if err != nil {
		c.JSON(401, res)
	}
	c.JSON(200, res)
}

func (h RequestHandler) SendEmailRegister(c *gin.Context) {
	email := c.Param("email")
	res, err := h.Ctrl.SendEmailRegister(email)
	if err != nil {
		c.JSON(401, res)
	}
	c.JSON(200, res)
}
func (h RequestHandler) CompareVerificationCode(c *gin.Context) {
	var verificationCodeRequest VerificationCodeRequest

	if err := c.ShouldBindJSON(&verificationCodeRequest); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
		return
	}
	res, err := h.Ctrl.CompareVerificationCode(&verificationCodeRequest)
	if err != nil {
		c.JSON(400, res)
		return
	}
	c.JSON(200, res)
}
func (h RequestHandler) EditPassword(c *gin.Context) {
	var req EditDataUserRequest
	queryUrl := c.Request.URL.Query()
	id := queryUrl.Get("id")
	if id == "" {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Bad Request Id Parameter kosong",
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"status":  "Error",
			"message": "Internal Server Error",
		})
		return
	}
	if req.Id == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request id nil",
		})
		return
	}

	code := c.Request.Header.Get("VerificationCode")
	if code == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"status":  "Error",
			"message": "Bad Request Verification Code nil",
		})
		return
	}

	res, err := h.Ctrl.EditPassword(id, code, &req)
	if err != nil {
		c.JSON(500, res)
	}
	c.JSON(200, res)
}

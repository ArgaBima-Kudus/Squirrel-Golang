package transaction

import (
	"golang/module/transaction/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandler struct {
	Ctrl ControllerInterface
}

type RequestHandlerinterface interface {
	GetAllTransactionByRequest(c *gin.Context)
}

func NewRequestHandler(ctrl ControllerInterface) RequestHandlerinterface {
	return RequestHandler{
		Ctrl: ctrl,
	}
}

func DefaultRequestHandler(db *gorm.DB) RequestHandlerinterface {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}

func (h RequestHandler) GetAllTransactionByRequest(c *gin.Context) {
	var req dto.Request

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch {
	case req.Status != "" || req.Status == "":
		if req.StartDate == "" && req.EndDate != "" {
			if req.Page == 0 && req.Size != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Page is required"})
				return
			} else if req.Page != 0 && req.Size == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Size Date is required"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Start Date is required"})
			return
		} else if req.StartDate != "" && req.EndDate == "" {
			if req.Page == 0 && req.Size != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Page is required"})
				return
			} else if req.Page != 0 && req.Size == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Size Date is required"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "End Date is required"})
			return
		}else {
			if req.Page == 0 && req.Size != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Page is required"})
				return
			} else if req.Page != 0 && req.Size == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Size Date is required"})
				return
			}
		}
	}

	transactions, err, _ := h.Ctrl.GetAllTransactionByRequest(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

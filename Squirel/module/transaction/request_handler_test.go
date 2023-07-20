package transaction

import (
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/module/transaction/dto"
	mocks "golang/module/transaction/mocks"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler_GetAllTransactionByRequest(t *testing.T) {
	
	type field struct {
		Ctrl ControllerInterface
	}

	tests := []struct {
		name string
		expectedStatusCode int
		makeRequest func() *http.Request
		makeFields func() field
		assertValue assert.ValueAssertionFunc
	}{
		{
			name: "Test Case Error",
			expectedStatusCode: http.StatusInternalServerError,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/get-transaction-by-status", nil)
				return req
			},
			makeFields: func() field {
				mockController := mocks.NewController(t)
				err := errors.New("error")
				mockController.EXPECT().GetAllTransactionByRequest(&dto.Request{}).Return(&dto.GetAllResponseDataTransaction{}, err, 0).Once()
				return field{
					Ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := gin.H{"error":"error", "message":"error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"error":"error", "message":"error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine)  {
				router.GET("/get-transaction-by-status", h.GetAllTransactionByRequest)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body){
				t.Errorf("assert value %v", body)
			}
			})
		}
	}


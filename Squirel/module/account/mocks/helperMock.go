package mocks

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewMockQueryDb creates a new mock database connection for testing purposes
// using sqlmock and gorm. If there is an error when creating the mock database,
// the function will call t.Fatal to stop the test and report the error.
// Example usage:
//
//	mockQuery, mockDb := test.NewMockQueryDb(t)
func NewMockQueryDb(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	conn, mockQuery, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(fmt.Errorf("sqlmock.New: %w", err))
	}
	mysqlConfig := mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}
	options := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Time{}
		},
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	mockDb, err := gorm.Open(mysql.New(mysqlConfig), options)
	if err != nil {
		t.Fatal(fmt.Errorf("gorm.Open: %w", err))
	}
	return mockQuery, mockDb
}

func CreateTestServer(request *http.Request, route func(router *gin.Engine)) (int, []byte) {
	gin.SetMode(gin.TestMode)
	responseRecorder := httptest.NewRecorder()
	_, router := gin.CreateTestContext(responseRecorder)

	route(router)
	router.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := io.ReadAll(response.Body)
	defer func() {
		_ = response.Body.Close()
	}()

	return response.StatusCode, body
}

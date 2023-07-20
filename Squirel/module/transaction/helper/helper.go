package helper

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
)

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

func NewMockQueryDb(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	conn, mockQuery, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal(fmt.Errorf("sqlmock.New: %w", err))
	}

	mysqlConfig := mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}
	options := gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}
	mockDb, _ := gorm.Open(mysql.New(mysqlConfig), &options)

	return mockQuery, mockDb
}

func FormatDate(date string) string {
	day := string(date[:2])
	month := string(date[3:5])
	year := string(date[6:10])

	parseDate := year + "-" + month + "-" + day
	return parseDate
}

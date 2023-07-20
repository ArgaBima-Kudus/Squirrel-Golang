package test

import (
	"golang/module/exportcsv"
	mocks "golang/module/exportcsv/test/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportCSVController(t *testing.T) {
	mockUsecase := mocks.NewUsecase(t)
	mockUsecase.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"", "", ""}).Return([][]string{{"id", "oda_number", "status", "price", "created_at"},
		{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
		{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
		{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
		{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}}, nil).Once()
	mockUsecase.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"SUCCESS", "", ""}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()

	type test struct {
		name    string
		request exportcsv.ExportCSVRequest
		url     exportcsv.UrlRequest
		expect  [][]string
	}

	tests := []test{
		test{
			name:    "PositifTest-exportAll",
			request: exportcsv.ExportCSVRequest{"", "", ""},
			url:     exportcsv.UrlRequest{"/export-transaction", "/export-transaction"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
		test{
			name:    "PositifTest-rangeDate filter",
			request: exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"},
			url:     exportcsv.UrlRequest{"/export-transaction?start_date=2023-04-17&end_date=2023-04-18", "/export-transaction"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
		test{
			name:    "PositifTest-SUCCESS status filter",
			request: exportcsv.ExportCSVRequest{"SUCCESS", "", ""},
			url:     exportcsv.UrlRequest{"/export-transaction?status=SUCCESS", "/export-transaction"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		},
		test{
			name:    "PositifTest-SUCCESS statusAndRageDate filter",
			request: exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"},
			url:     exportcsv.UrlRequest{"/export-transaction?status=SUCCESS&start_date=2023-04-17&end_date=2023-04-18", "/export-transaction"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		},
		test{
			name:    "PositifTest-WAITING_FOR_DEBITTED filter",
			request: exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""},
			url:     exportcsv.UrlRequest{"/export-transaction?status=WAITING_FOR_DEBITTED", "/export-transaction"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
		test{
			name:    "Positif-WAITING_FOR_DEBITTTED statusAndRageDate filter",
			request: exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"},
			url:     exportcsv.UrlRequest{"/export-transaction?status=WAITING_FOR_DEBITTED&start_date=2023-04-17&end_date=2023-04-18", "/export-transaction"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
	}
	controller := exportcsv.NewController(mockUsecase)
	for _, forTest := range tests {
		t.Run(forTest.name, func(t *testing.T) {
			result, err := controller.ExportCSV(&forTest.request, &forTest.url)
			assert.Nil(t, err)
			assert.Equal(t, forTest.expect, result)

		})
	}

}

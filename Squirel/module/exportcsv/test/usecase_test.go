package test

import (
	"golang/module/exportcsv"
	mocks "golang/module/exportcsv/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionStringConverter(t *testing.T) {
	request1 := []exportcsv.Transaction{{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"}}
	request2 := []exportcsv.Transaction{{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"}}
	request3 := []exportcsv.Transaction{{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"}}
	request4 := []exportcsv.Transaction{{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"}}
	expected1 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"}}
	expected2 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}}
	expected3 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}}
	expected4 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}}
	tests := []struct {
		name     string
		request  []exportcsv.Transaction
		expected [][]string
	}{
		{"positif1", request1, expected1},
		{"positif2", request2, expected2},
		{"positif3", request3, expected3},
		{"positif4", request4, expected4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := exportcsv.TransactionStringConverter(test.request)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, nil, err)

		})
	}

}
func TestGet(t *testing.T) {
	type args struct {
		req *exportcsv.ExportCSVRequest
	}
	request1 := []exportcsv.Transaction{{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"}}
	mockUsecase := mocks.NewUsecase(t)
	mockUsecase.EXPECT().Get().Return(request1, nil).Once()

	got, err := mockUsecase.Get()
	assert.Nil(t, err)
	assert.Equal(t, request1, got)

}
func TestExportCSV(t *testing.T) {

	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().GetAllTransaction().Return([]exportcsv.Transaction{
		{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetAllTransactionByRangeDateFilter(&exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"}).Return([]exportcsv.Transaction{
		{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusFilter(&exportcsv.ExportCSVRequest{"SUCCESS", "", ""}).Return([]exportcsv.Transaction{
		{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusAndRangeDateFilter(&exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"}).Return([]exportcsv.Transaction{
		{2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusFilter(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""}).Return([]exportcsv.Transaction{
		{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusAndRangeDateFilter(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"}).Return([]exportcsv.Transaction{
		{241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
		{242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17"},
	}, nil).Once()

	type test struct {
		name    string
		request exportcsv.ExportCSVRequest
		expect  [][]string
	}

	tests := []test{
		test{
			name:    "PositifTest-exportAll",
			request: exportcsv.ExportCSVRequest{"", "", ""},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		}, test{
			name:    "PositifTest-rangeDate filter",
			request: exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		}, test{
			name:    "PositifTest-SUCCESS status filter",
			request: exportcsv.ExportCSVRequest{"SUCCESS", "", ""},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		}, test{
			name:    "PositifTest-SUCCESS statusAndRageDate filter",
			request: exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		}, test{
			name:    "PositifTest-WAITING_FOR_DEBITTED filter",
			request: exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		}, test{
			name:    "Positif-WAITING_FOR_DEBITTTED statusAndRageDate filter",
			request: exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
	}

	usecase := exportcsv.NewUseCase(mockRepo)
	for _, forTest := range tests {
		t.Run(forTest.name, func(t *testing.T) {
			result, err := usecase.ExportCSV(&forTest.request)
			assert.Nil(t, err)
			assert.Equal(t, forTest.expect, result)

		})
	}

}

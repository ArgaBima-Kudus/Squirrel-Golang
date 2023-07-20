package exportcsv

import (
	"errors"
)

type controller struct {
	useCase Usecase
}

func NewController(useCase Usecase) *controller {
	return &controller{
		useCase: useCase,
	}
}

type ExportCSVResponse struct {
	ExportCSVResponseMasage string `json:"masage"`
}

type TransactionCollection struct {
	DataItem []Transaction
}
type Controller interface {
	ExportCSV(req *ExportCSVRequest, url *UrlRequest) ([][]string, error)
}

func (c controller) ExportCSV(req *ExportCSVRequest, url *UrlRequest) ([][]string, error) {
	isRequestInvaid := false
	var clientWant = url.ClientUrl
	var serverGot string
	switch {
	case (req.StartDate == "" && req.EndDate != "") || (req.StartDate != "" && req.EndDate == ""):
		isRequestInvaid = true
	case req.StartDate != "" && req.StartDate == "" && req.EndDate == "":
		serverGot = url.PathUrl + "?" + "status=" + req.Status
		if clientWant != serverGot {
			isRequestInvaid = true
		}

	case req.Status == "" && req.StartDate != "" && req.EndDate != "":
		serverGot = url.PathUrl + "?" + "start_date=" + req.StartDate + "&end_date=" + req.EndDate
		if serverGot != clientWant {
			isRequestInvaid = true
		}
	case req.Status != "" && req.StartDate != "" && req.EndDate != "":
		serverGot = url.PathUrl + "?" + "status=" + req.Status + "&start_date=" + req.StartDate + "&end_date=" + req.EndDate
		if serverGot != clientWant {
			isRequestInvaid = true
		}
	}
	if isRequestInvaid == false {
		switch req.Status {
		case "SUCCESS", "WAITING_FOR_DEBITTED", "": //validate request.status
			exportData, err := c.useCase.ExportCSV(req)
			if err != nil {

				return nil, err
			}
			return exportData, nil

		default:
			return nil, errors.New("Invalid Request")
		}
	}
	return nil, errors.New("Invalid Request")
}

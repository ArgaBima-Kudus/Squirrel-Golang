package dto

type TransactionItemResponse struct {
	Id         int     `json:"id"`
	OdaNumber  string  `json:"oda_number"`
	BillAmount float32 `json:"price"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type GetAllResponseDataTransaction struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Error     string                    `json:"status"`
	TotalData int                       `json:"total_data"`
	TotalDataFiltered int                `json:"total_data_filtered"`
	Data      []TransactionItemResponse `json:"data"`
}

type Request struct {
	Status    string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

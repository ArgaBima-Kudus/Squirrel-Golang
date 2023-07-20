package exportcsv

type Transaction struct {
	Id                 int
	Oda_number         string
	Bank_account_no    string
	Billing_cycle_date string
	Payment_due_date   string
	Overflow_amount    float64
	Bill_amount        float64
	Principal_amount   float64
	Interest_amount    float64
	Total_fee_amount   float64
	Status             string
	Payment_method     string
	Auto_debet_counter int
	Created_at         string
	Updated_at         string
	Payment_amount     float64
	Billing_gen_date   string
}

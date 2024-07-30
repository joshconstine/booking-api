package response

type TaxRateResponse struct {
	ID         uint    `json:"id"`
	Percentage float64 `json:"percentage"`
	Name       string  `json:"name"`
}

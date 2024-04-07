package request

type CreateTaxRateRequest struct {
	Percentage float64 `json:"percentage"`
	Name       string  `json:"name"`
}

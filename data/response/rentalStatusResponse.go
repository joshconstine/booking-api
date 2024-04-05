package response

type RentalStatusResponse struct {
	RentalID uint `json:"rentalId"`
	IsClean  bool `json:"isClean"`
}

package request

type UpdateRentalStatusRequest struct {
	IsClean  bool `json:"isClean"`
	RentalID uint `json:"rentalId"`
}

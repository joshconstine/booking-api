package response

type BoatStatusResponse struct {
	ID         uint `json:"id"`
	BoatID     uint `json:"boatID"`
	IsClean    bool `json:"isClean"`
	LocationID uint `json:"locationID"`
}

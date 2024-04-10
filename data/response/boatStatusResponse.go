package response

type BoatStatusResponse struct {
	BoatID     uint `json:"boatID"`
	IsClean    bool `json:"isClean"`
	LocationID uint `json:"locationID"`
}

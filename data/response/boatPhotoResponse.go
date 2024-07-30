package response

type BoatPhotoResponse struct {
	ID       uint   `json:"id"`
	BoatID   int    `json:"boatID"`
	PhotoURL string `json:"photoURL"`
}

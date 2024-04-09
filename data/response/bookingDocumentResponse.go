package response

type BookingDocumentResponse struct {
	ID                uint             `json:"id"`
	BookingID         string           `json:"booking_id"`
	RequiresSignature bool             `json:"requires_signature"`
	Signed            bool             `json:"signed"`
	Note              string           `json:"note"`
	Document          DocumentResponse `json:"document"`
}

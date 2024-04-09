package request

type CreateEntityBookingDocumentRequest struct {
	EntityID          uint
	EntityType        string
	DocumentID        uint
	RequiresSignature bool
}

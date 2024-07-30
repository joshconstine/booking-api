package response

type EntityBookingDocumentResponse struct {
	ID                uint
	RequiresSignature bool
	Document          DocumentResponse
}

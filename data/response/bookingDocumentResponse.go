package response

type BookingDocumentResponse struct {
	ID                uint
	RequiresSignature bool
	Signed            bool
	Note              string
	Document          DocumentResponse
}

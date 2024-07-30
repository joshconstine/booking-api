package request

type GetBookingSnapshotRequest struct {
	SearchString      string            `json:"searchString"`
	Statuses          []int             `json:"statuses"`
	PaginationRequest PaginationRequest `json:"paginationRequest"`
}

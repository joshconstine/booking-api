package response

type EntityInfoResponse struct {
	EntityID   uint   `json:"entityID"`
	EntityType string `json:"entityType"`
	Name       string `json:"name"`
}

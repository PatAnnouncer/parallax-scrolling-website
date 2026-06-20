package models

// ServerStatus represents the response structure for the frontend
type ServerStatus struct {
	Minecraft bool `json:"minecraft"`
	DnD       bool `json:"dnd"`
}

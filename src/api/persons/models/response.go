package models

type ErrorResponse struct {
	Message string   `json:"message,omitempty"`
	Details []string `json:"details,omitempty"`
}

package account

import uuid "github.com/satori/go.uuid"

// CreatePayload create account request
type CreatePayload struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedBy uuid.UUID
}

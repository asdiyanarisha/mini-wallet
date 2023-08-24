package model

import "github.com/google/uuid"

type (
	Transaction struct {
		Id            uuid.UUID `json:"id"`
		Status        string    `json:"status"`
		TransactionAt string    `json:"transaction_at"`
		Type          string    `json:"type"`
		Amount        int       `json:"amount"`
		ReferenceId   uuid.UUID `json:"reference_id"`
	}
)

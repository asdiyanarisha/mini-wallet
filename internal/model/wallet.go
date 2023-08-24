package model

import "github.com/google/uuid"

type (
	Wallet struct {
		Id          uuid.UUID `json:"id"`
		CustomerXid uuid.UUID `json:"customer_xid"`
		Status      string    `json:"status"`
		Token       string    `json:"token,omitempty"`
		EnabledAt   string    `json:"enabled_at"`
		Balance     int       `json:"balance"`
	}
)

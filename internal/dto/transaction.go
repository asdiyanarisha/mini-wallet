package dto

import "github.com/google/uuid"

type (
	DepositWallet struct {
		Amount      string `form:"amount" binding:"required"`
		ReferenceId string `form:"reference_id" binding:"required"`
	}

	WithdrawalWallet struct {
		Amount      string `form:"amount" binding:"required"`
		ReferenceId string `form:"reference_id" binding:"required"`
	}

	ResponseDeposit struct {
		Id          uuid.UUID `json:"id"`
		DepositedBy uuid.UUID `json:"deposited_by"`
		Status      string    `json:"status"`
		DepositedAt string    `json:"deposited_at"`
		Amount      int       `json:"amount"`
		ReferenceId uuid.UUID `json:"reference_id"`
	}

	ResponseWithdrawal struct {
		Id          uuid.UUID `json:"id"`
		WithdrawnBy uuid.UUID `json:"withdrawn_by"`
		Status      string    `json:"status"`
		WithdrawnAt string    `json:"withdrawn_at"`
		Amount      int       `json:"amount"`
		ReferenceId uuid.UUID `json:"reference_id"`
	}

	ResponseDepositInit struct {
		Deposit ResponseDeposit `json:"deposit"`
	}

	ResponseWithdrawalInit struct {
		Withdrawal ResponseWithdrawal `json:"withdrawal"`
	}
)

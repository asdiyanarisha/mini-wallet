package dto

type (
	DepositWallet struct {
		Amount      string `form:"amount" binding:"required"`
		ReferenceId string `form:"reference_id" binding:"required"`
	}
)

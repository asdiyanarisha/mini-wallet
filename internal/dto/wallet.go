package dto

type (
	InitializeWallet struct {
		CustomerXid string `form:"customer_xid" binding:"required"`
	}
)

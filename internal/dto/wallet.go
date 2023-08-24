package dto

type (
	InitializeWallet struct {
		CustomerXid string `form:"customer_xid" binding:"required"`
	}

	ResponseDataEnable struct {
		Id        string `json:"id"`
		OwnedBy   string `json:"owned_by"`
		Status    string `json:"status"`
		EnabledAt string `json:"enabled_at"`
		Balance   int    `json:"balance"`
	}

	ResponseWalletEnabled struct {
		Wallet ResponseDataEnable `json:"wallet"`
	}
)

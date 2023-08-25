package dto

type (
	InitializeWallet struct {
		CustomerXid string `form:"customer_xid" binding:"required"`
	}

	DisableWallet struct {
		IsDisabled bool `form:"is_disabled" binding:"required"`
	}

	ResponseDataEnable struct {
		Id        string `json:"id"`
		OwnedBy   string `json:"owned_by"`
		Status    string `json:"status"`
		EnabledAt string `json:"enabled_at"`
		Balance   int    `json:"balance"`
	}

	ResponseDataDisable struct {
		Id         string `json:"id"`
		OwnedBy    string `json:"owned_by"`
		Status     string `json:"status"`
		DisabledAt string `json:"disabled_at"`
		Balance    int    `json:"balance"`
	}

	ResponseWallet struct {
		Wallet any `json:"wallet"`
	}
)

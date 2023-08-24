package dto

type (
	Error struct {
		Error any `json:"error"`
	}

	Common struct {
		Status string `json:"status"`
		Data   any    `json:"data"`
	}

	ResponseToken struct {
		Token string `json:"token"`
	}
)

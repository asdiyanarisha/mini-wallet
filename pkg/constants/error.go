package constants

import "errors"

var (
	InternalServerError       = errors.New("internal server error")
	AlreadyEnabledError       = errors.New("already enabled")
	WalletDisabledError       = errors.New("wallet disabled")
	ReferenceIdAlreadyTracked = errors.New("reference id already usage")
)

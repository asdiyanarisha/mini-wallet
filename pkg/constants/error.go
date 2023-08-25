package constants

import "errors"

var (
	InternalServerError       = errors.New("internal server error")
	AlreadyEnabledError       = errors.New("already enabled")
	WalletDisabledError       = errors.New("wallet disabled")
	IsDisabledMustTrueError   = errors.New("parameter is_disabled must be true")
	InsufficientBalanceError  = errors.New("insufficient balance")
	ReferenceIdAlreadyTracked = errors.New("reference id already usage")
)

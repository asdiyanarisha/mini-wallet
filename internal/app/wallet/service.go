package wallet

import (
	"github.com/google/uuid"
	"mini-wallet/internal/dto"
	"mini-wallet/internal/model"
	"mini-wallet/pkg/constants"
	"mini-wallet/pkg/helper"
)

type service struct {
}

type Service interface {
	InitializeData(payload dto.InitializeWallet) (string, error)
	EnableWallet(wallet model.Wallet) (dto.ResponseWallet, error)
	GetBalanceWallet(wallet model.Wallet) (dto.ResponseWallet, error)
	DisableWallet(wallet model.Wallet, payload dto.DisableWallet) (dto.ResponseWallet, error)
}

func NewService() Service {
	return &service{}
}

func (s *service) GetBalanceWallet(wallet model.Wallet) (dto.ResponseWallet, error) {
	if wallet.Status != "enabled" {
		return dto.ResponseWallet{}, constants.WalletDisabledError
	}

	response := dto.ResponseDataEnable{
		Id:        wallet.Id.String(),
		OwnedBy:   wallet.CustomerXid.String(),
		Status:    wallet.Status,
		EnabledAt: wallet.EnabledAt,
		Balance:   wallet.Balance,
	}

	return dto.ResponseWallet{Wallet: response}, nil
}

func (s *service) DisableWallet(wallet model.Wallet, payload dto.DisableWallet) (dto.ResponseWallet, error) {
	if wallet.Status == "disabled" {
		return dto.ResponseWallet{}, constants.WalletDisabledError
	}

	if payload.IsDisabled != true {
		return dto.ResponseWallet{}, constants.IsDisabledMustTrueError
	}

	wallet.Status = "disabled"

	helper.WriteJson(wallet, wallet.CustomerXid.String())

	response := dto.ResponseDataDisable{
		Id:         wallet.Id.String(),
		OwnedBy:    wallet.CustomerXid.String(),
		Status:     wallet.Status,
		DisabledAt: helper.InitDate(),
		Balance:    0,
	}

	return dto.ResponseWallet{Wallet: response}, nil
}

func (s *service) EnableWallet(wallet model.Wallet) (dto.ResponseWallet, error) {
	if wallet.Status == "enabled" {
		return dto.ResponseWallet{}, constants.AlreadyEnabledError
	}

	wallet.Status = "enabled"
	wallet.EnabledAt = helper.InitDate()

	helper.WriteJson(wallet, wallet.CustomerXid.String())

	response := dto.ResponseDataEnable{
		Id:        wallet.Id.String(),
		OwnedBy:   wallet.CustomerXid.String(),
		Status:    wallet.Status,
		EnabledAt: wallet.EnabledAt,
		Balance:   wallet.Balance,
	}

	return dto.ResponseWallet{Wallet: response}, nil
}

func (s *service) InitializeData(payload dto.InitializeWallet) (string, error) {
	// get all wallets if file not already created, this app return empty array
	wallet, err := helper.OpenWalletFile(payload.CustomerXid)
	if err != nil {
		return "", err
	}

	if wallet == (model.Wallet{}) {
		// if wallet has enabled, app will return by customer idx
		token, _ := helper.RandomHex(20)

		customerXid, err := uuid.Parse(payload.CustomerXid)
		if err != nil {
			return "", err
		}

		// build format data wallet
		data := model.Wallet{
			Id:          helper.GetUuid(),
			CustomerXid: customerXid,
			Status:      "disabled",
			Token:       token,
			Balance:     0,
		}

		if err := helper.UpdateToken(token, data.CustomerXid.String()); err != nil {
			return "", err
		}

		helper.WriteJson(data, data.CustomerXid.String())
		return data.Token, nil
	} else {
		return wallet.Token, nil
	}
}

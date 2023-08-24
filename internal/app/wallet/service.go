package wallet

import (
	"github.com/google/uuid"
	"julo-test/internal/dto"
	"julo-test/internal/model"
	"julo-test/pkg/helper"
)

type service struct {
}

type Service interface {
	InitializeData(payload dto.InitializeWallet) (string, error)
}

func NewService() Service {
	return &service{}
}

func (s *service) InitializeData(payload dto.InitializeWallet) (string, error) {
	// get all wallets if file not already created, this app return empty array
	wallets, err := helper.OpenWalletFile()
	if err != nil {
		return "", err
	}

	// filter wallets array by customer idx
	wallet, err := s.GetWalletByCustId(wallets, payload.CustomerXid)
	if err != nil {
		return "", err
	}

	if wallet != (model.Wallet{}) {
		// if wallet has enabled, app will return by customer idx
		return wallet.Token, nil
	}

	token, _ := helper.RandomHex(20)

	customerXid, err := uuid.Parse(payload.CustomerXid)
	if err != nil {
		return "", err
	}

	// build format data wallet
	data := model.Wallet{
		Id:          helper.GetUuid(),
		CustomerXid: customerXid,
		Status:      "enabled",
		Token:       token,
		EnabledAt:   helper.InitDate(),
		Balance:     0,
	}

	wallets = append(wallets, data)
	helper.WriteJson(wallets)

	return data.Token, nil
}

func (s *service) GetWalletByCustId(wallets []model.Wallet, customerIdx string) (model.Wallet, error) {
	for _, wallet := range wallets {
		if wallet.CustomerXid.String() == customerIdx && wallet.Status == "enabled" {
			return wallet, nil
		}
	}

	return model.Wallet{}, nil
}

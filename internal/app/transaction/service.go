package transaction

import (
	"julo-test/internal/dto"
	"julo-test/internal/model"
	"julo-test/pkg/helper"
)

type service struct {
}

type Service interface {
}

func NewService() Service {
	return &service{}
}

func (s *service) DepositWallet(myWallet model.Wallet, request dto.DepositWallet) error {
	_, err := helper.OpenTransactionFile(myWallet.CustomerXid.String())
	if err != nil {
		return err
	}

	return nil

}

package transaction

import (
	"github.com/google/uuid"
	"julo-test/internal/dto"
	"julo-test/internal/model"
	"julo-test/pkg/constants"
	"julo-test/pkg/helper"
	"strconv"
)

type service struct {
}

type Service interface {
	DepositWallet(myWallet model.Wallet, request dto.DepositWallet) (dto.ResponseDeposit, error)
}

func NewService() Service {
	return &service{}
}

func (s *service) DepositWallet(myWallet model.Wallet, request dto.DepositWallet) (dto.ResponseDeposit, error) {
	amount, _ := strconv.Atoi(request.Amount)
	referenceId, err := uuid.Parse(request.ReferenceId)
	if err != nil {
		return dto.ResponseDeposit{}, err
	}

	transactions, err := helper.OpenTransactionFile(myWallet.CustomerXid.String())
	if err != nil {
		return dto.ResponseDeposit{}, err
	}

	if err := s.CheckReferenceId(transactions, referenceId); err != nil {
		return dto.ResponseDeposit{}, err
	}

	transaction := model.Transaction{
		Id:            helper.GetUuid(),
		Status:        "success",
		TransactionAt: helper.InitDate(),
		Type:          "deposit",
		Amount:        amount,
		ReferenceId:   referenceId,
	}
	transactions = append(transactions, transaction)

	helper.WriteTransaction(transactions, myWallet.CustomerXid.String())

	// update balance
	myWallet.Balance = myWallet.Balance + amount
	s.UpdateBalance(myWallet)

	response := dto.ResponseDeposit{
		Id:          transaction.Id,
		DepositedBy: myWallet.CustomerXid,
		Status:      "success",
		DepositedAt: transaction.TransactionAt,
		Amount:      transaction.Amount,
		ReferenceId: referenceId,
	}
	return response, nil

}

func (s *service) UpdateBalance(myWallet model.Wallet) {
	helper.WriteJson(myWallet, myWallet.CustomerXid.String())
}

func (s *service) CheckReferenceId(transactions []model.Transaction, referenceId uuid.UUID) error {
	for _, transaction := range transactions {
		if transaction.ReferenceId == referenceId {
			return constants.ReferenceIdAlreadyTracked
		}
	}

	return nil
}

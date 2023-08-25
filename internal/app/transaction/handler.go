package transaction

import (
	"github.com/gin-gonic/gin"
	"julo-test/internal/dto"
	"julo-test/internal/model"
	"julo-test/pkg/helper"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler() *handler {
	return &handler{
		service: NewService(),
	}
}

func (h *handler) WithdrawalWalletHandler(g *gin.Context) {
	var payload dto.WithdrawalWallet
	if err := g.ShouldBind(&payload); err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: helper.Validate(err)},
		})
		return
	}
	wallet, _ := g.Get("wallet")

	response, err := h.service.WithdrawalWallet(wallet.(model.Wallet), payload)
	if err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: err.Error()},
		})
		return
	}

	g.JSON(http.StatusCreated, dto.Common{Status: "success", Data: dto.ResponseWithdrawalInit{Withdrawal: response}})
	return
}

func (h *handler) DepositWalletHandler(g *gin.Context) {
	var payload dto.DepositWallet
	if err := g.ShouldBind(&payload); err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: helper.Validate(err)},
		})
		return
	}
	wallet, _ := g.Get("wallet")

	response, err := h.service.DepositWallet(wallet.(model.Wallet), payload)
	if err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: err.Error()},
		})
		return
	}

	g.JSON(http.StatusCreated, dto.Common{Status: "success", Data: dto.ResponseDepositInit{Deposit: response}})
	return
}

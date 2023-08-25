package wallet

import (
	"github.com/gin-gonic/gin"
	"mini-wallet/internal/dto"
	"mini-wallet/internal/model"
	"mini-wallet/pkg/helper"
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

func (h *handler) GetBalanceHandler(g *gin.Context) {
	wallet, _ := g.Get("wallet")

	response, err := h.service.GetBalanceWallet(wallet.(model.Wallet))
	if err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: err.Error()},
		})
		return
	}

	g.JSON(http.StatusOK, dto.Common{Status: "success", Data: response})
	return
}

func (h *handler) DisableWalletHandler(g *gin.Context) {
	wallet, _ := g.Get("wallet")
	var payload dto.DisableWallet
	if err := g.ShouldBind(&payload); err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: helper.Validate(err)},
		})
		return
	}

	response, err := h.service.DisableWallet(wallet.(model.Wallet), payload)
	if err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: err.Error()},
		})
		return
	}

	g.JSON(http.StatusCreated, dto.Common{Status: "success", Data: response})
}

func (h *handler) EnableWalletHandler(g *gin.Context) {
	wallet, _ := g.Get("wallet")

	response, err := h.service.EnableWallet(wallet.(model.Wallet))
	if err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: err.Error()},
		})
		return
	}

	g.JSON(http.StatusCreated, dto.Common{Status: "success", Data: response})
	return
}

func (h *handler) InitializeHandler(g *gin.Context) {
	var payload dto.InitializeWallet
	if err := g.ShouldBind(&payload); err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: helper.Validate(err)},
		})
		return
	}

	token, err := h.service.InitializeData(payload)
	if err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: err.Error()},
		})
		return
	}

	g.JSON(http.StatusCreated, dto.Common{Status: "success", Data: dto.ResponseToken{Token: token}})
	return
}

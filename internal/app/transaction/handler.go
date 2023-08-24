package transaction

import (
	"github.com/gin-gonic/gin"
	"julo-test/internal/dto"
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

func (h *handler) DepositWalletHandler(g *gin.Context) {
	var payload dto.DepositWallet
	if err := g.ShouldBind(&payload); err != nil {
		g.JSON(http.StatusBadRequest, dto.Common{
			Status: "fail",
			Data:   dto.Error{Error: helper.Validate(err)},
		})
		return
	}

}
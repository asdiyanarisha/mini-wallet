package wallet

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

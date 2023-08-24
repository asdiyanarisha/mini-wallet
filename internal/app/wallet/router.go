package wallet

import "github.com/gin-gonic/gin"

func (h *handler) WalletRouter(g *gin.RouterGroup) {
	g.POST("", h.EnableWalletHandler)
	g.GET("", h.GetBalanceHandler)
}

func (h *handler) InitWallet(g *gin.RouterGroup) {
	g.POST("init", h.InitializeHandler)
}

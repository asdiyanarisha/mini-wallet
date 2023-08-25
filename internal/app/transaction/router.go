package transaction

import "github.com/gin-gonic/gin"

func (h *handler) TransactionRouter(g *gin.RouterGroup) {
	g.POST("deposits", h.DepositWalletHandler)
	g.POST("withdrawals", h.WithdrawalWalletHandler)
	g.GET("transactions", h.TransactionWalletHandler)
}

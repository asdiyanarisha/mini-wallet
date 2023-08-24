package http

import (
	"github.com/gin-gonic/gin"
	"julo-test/internal/app/wallet"
	"julo-test/internal/middleware"
)

func NewHttp(g *gin.Engine) {
	g.Use(middleware.CORSMiddleware(), gin.Logger(), gin.Recovery())
	v1 := g.Group("/api/v1")

	wallet.NewHandler().InitWallet(v1)

	walletGroup := v1.Group("/wallet")

	walletGroup.Use(middleware.BearerToken())
	{
		wallet.NewHandler().WalletRouter(walletGroup)
	}
}

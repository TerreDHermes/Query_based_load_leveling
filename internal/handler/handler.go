package handler

import (
	"backend/internal/service"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/api", h.Hello)
	router.GET("/create", h.CreateWallet)
	router.GET("/info/:walletId", h.WalletInfo)
	return router
}

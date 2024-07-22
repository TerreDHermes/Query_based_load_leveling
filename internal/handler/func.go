package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}

func (h *Handler) CreateWallet(c *gin.Context) {
	id, err := h.service.CreateWallet()
	if err != nil {
		logrus.Error("Кошелек не создан", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func (h *Handler) WalletInfo(c *gin.Context) {
	wallet := c.Param("walletId")

	walletId, err := strconv.Atoi(wallet)
	if err != nil {
		// Если параметр не является числом, возвращаем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
	}
	id, balance, err := h.service.WalletInfo(walletId)
	if err != nil {
		logrus.Error("Кошелек не найден", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":      id,
			"balance": balance,
		})
	}
}

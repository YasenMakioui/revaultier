package card

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CardHandler struct {
	CardService *CardService
}

func NewCardHandler(cs *CardService) *CardHandler {
	return &CardHandler{CardService: cs}
}

func (h *CardHandler) GetCardHandler(c echo.Context) error {
	return c.String(http.StatusOK, "card")
}

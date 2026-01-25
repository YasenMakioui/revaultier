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

	ctx := c.Request().Context()

	vaultId := c.Param("id")
	cardId := c.Param("cardId")

	card, err := h.CardService.GetCardService(ctx, vaultId, cardId)

	if err != nil || card.Id == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "card not found"})
	}

	return c.JSON(http.StatusOK, card)
}

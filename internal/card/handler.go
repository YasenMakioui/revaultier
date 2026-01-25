package card

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

	userId, err := getUserId(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no subject found in claims"})
	}

	vaultId := c.Param("id")
	cardId := c.Param("cardId")

	card, err := h.CardService.GetCardService(ctx, vaultId, cardId, userId)

	if err != nil || card.Id == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "card not found"})
	}

	return c.JSON(http.StatusOK, card)
}

// Helpers

func getUserId(c echo.Context) (string, error) {

	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	userId, err := claims.GetSubject()

	if err != nil {
		return "", err
	}

	return userId, nil
}

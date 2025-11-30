package vault

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type VaultHandler struct {
	VaultService *VaultService
}

func NewVaultHandler(vs *VaultService) *VaultHandler {
	return &VaultHandler{VaultService: vs}
}

func (h *VaultHandler) GetVaultsHandler(c echo.Context) error {

	ctx := c.Request().Context()

	jwtToken := c.Request().Header.Get("Authorization")

	return c.String(http.StatusOK, "test")
}

package vault

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	uuid, err := claims.GetSubject()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no subject found in claims"})
	}

	vaults, err := h.VaultService.GetVaultsService(ctx, uuid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, vaults)
}

func (h *VaultHandler) GetVaultHandler(c echo.Context) error {

	ctx := c.Request().Context()

	// maybe move the token thing to a helper function to get the subject

	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	userId, err := claims.GetSubject()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no subject found in claims"})
	}

	vaultId := c.Param("id")

	vault, err := h.VaultService.GetVaultService(ctx, vaultId, userId)

	if err != nil || vault.Id == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "vault not found"})
	}

	return c.JSON(http.StatusOK, vault)
}

func (h *VaultHandler) CreateVaultHandler(c echo.Context) error {

	//ctx := c.Request().Context()

	//token := Get("user").(*jwt.Token)

	//claims := token.Claims.(jwt.MapClaims)

	//userId, err := claims.GetSubject()

	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no subject found in claims"})
	//}

	// get params from payload and validate
	//vault, err := h.VaultService.CreateVaultService()
	return nil
}

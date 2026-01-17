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

	userId, err := getUserId(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no subject found in claims"})
	}

	vaults, err := h.VaultService.GetVaultsService(ctx, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "something went wrong"})
	}

	return c.JSON(http.StatusOK, vaults)
}

func (h *VaultHandler) GetVaultHandler(c echo.Context) error {

	ctx := c.Request().Context()

	userId, err := getUserId(c)

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

	ctx := c.Request().Context()

	userId, err := getUserId(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no subject found in claims"})
	}

	v := new(VaultDTO)

	if err := c.Bind(v); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "bad request"})
	}

	vault, err := h.VaultService.CreateVaultService(ctx, userId, v)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "bad request"})
	}

	return c.JSON(http.StatusOK, vault)
}

func (h *VaultHandler) DeleteVaultHandler(c echo.Context) error {

	ctx := c.Request().Context()

	vaultId := c.Param("id")

	if err := h.VaultService.DeleteVaultService(ctx, vaultId); err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusOK)
}

func getUserId(c echo.Context) (string, error) {

	token := c.Get("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	userId, err := claims.GetSubject()

	if err != nil {
		return "", err
	}

	return userId, nil
}

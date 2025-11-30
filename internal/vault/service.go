package vault

import (
	"context"
	"database/sql"
	"errors"
	"revaultier/configuration"
)

type VaultService struct {
	Cfg             *configuration.Config
	VaultRepository *VaultRepository
}

func NewVaultService(cfg *configuration.Config, vr *VaultRepository) *VaultService {
	return &VaultService{Cfg: cfg, VaultRepository: vr}
}

func (s *VaultService) GetVaultService(ctx context.Context, vaultId string, ownerId string) (Vault, error) {

	vault, err := s.VaultRepository.GetVault(ctx, vaultId, ownerId)

	if err != nil {
		if err == sql.ErrNoRows {
			return Vault{}, nil
		}
		return Vault{}, errors.New("could not get vault")
	}

	return vault, nil
}

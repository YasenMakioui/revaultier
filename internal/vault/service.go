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

func (s *VaultService) GetVaultsService(ctx context.Context, ownerId string) ([]Vault, error) {
	vaults, err := s.VaultRepository.GetVaults(ctx, ownerId)

	if err != nil {
		return vaults, err
	}

	return vaults, err
}

func (s *VaultService) CreateVaultService(ctx context.Context, ownerId string, name string, description string, created_at string) (Vault, error) {
	vault, err := s.VaultRepository.InsertVault(ctx, ownerId, name, description, created_at)

	if err != nil {
		return vault, err
	}

	return vault, nil
}

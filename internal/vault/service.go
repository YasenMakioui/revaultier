package vault

import (
	"context"
	"database/sql"
	"errors"
	"revaultier/configuration"
	"time"

	"github.com/google/uuid"
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

func (s *VaultService) CreateVaultService(ctx context.Context, ownerId string, v *VaultDTO) (Vault, error) {

	if v.Name == "" {
		return Vault{}, errors.New("name is empty")
	}

	vaultId := uuid.New().String()

	createdAt := time.Now().Format("2006-01-02")

	vault := Vault{
		Id:          vaultId,
		Owner_id:    ownerId,
		Name:        v.Name,
		Description: v.Description,
		Created_at:  createdAt,
	}

	vault, err := s.VaultRepository.InsertVault(ctx, vault)

	if err != nil {
		return vault, err
	}

	return vault, nil
}

func (s *VaultService) DeleteVaultService(ctx context.Context, vaultId string, ownerId string) error {

	if err := s.VaultRepository.DeleteVault(ctx, vaultId, ownerId); err != nil {
		return err
	}

	return nil
}

func (s *VaultService) UpdateVaultService(ctx context.Context, v *VaultDTO, vaultId string, ownerId string) error {

	if v.Name == "" {
		return errors.New("name is empty")
	}

	if err := s.VaultRepository.UpdateVault(ctx, v.Name, v.Description, vaultId, ownerId); err != nil {
		return err
	}

	return nil
}

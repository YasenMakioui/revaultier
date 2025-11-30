package vault

import (
	"context"
	"database/sql"
)

type VaultRepository struct {
	dbconn *sql.DB
}

func NewVaultRepository(db *sql.DB) *VaultRepository {
	return &VaultRepository{dbconn: db}
}

func (r *VaultRepository) GetVault(ctx context.Context, vaultId string, ownerId string) (Vault, error) {

	sqlStmt := "SELECT * FROM vault WHERE id=?"

	var vault Vault

	err := r.dbconn.QueryRowContext(ctx, sqlStmt, vaultId).Scan(&vault.Id, &vault.Owner_id, &vault.Name, &vault.Description, &vault.Created_at)

	if err != nil {
		return Vault{}, err
	}

	return vault, err
}

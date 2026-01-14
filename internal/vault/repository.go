package vault

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
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

func (r *VaultRepository) GetVaults(ctx context.Context, ownerId string) ([]Vault, error) {

	sqlStmt := "SELECT * FROM vault WHERE owner_id=?"

	var vaults []Vault

	rows, err := r.dbconn.QueryContext(ctx, sqlStmt, ownerId)

	for rows.Next() {

		if err != nil {
			break
		}

		var vault Vault

		err := rows.Scan(&vault.Id, &vault.Owner_id, &vault.Name, &vault.Description, &vault.Created_at)

		if err != nil {
			return vaults, err
		}

		vaults = append(vaults, vault)
	}

	return vaults, nil
}

func (r *VaultRepository) InsertVault(ctx context.Context, ownerId string, name string, description string, created_at string) (Vault, error) {
	sqlStmt := "INSERT INTO vault(id,owner_id,name,description,created_at) values(?,?,?,?,?)"

	vaultId := uuid.NewString()

	var vault Vault

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, vaultId, ownerId, name, description, created_at)

	if err != nil {
		return vault, err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return vault, errors.New("could not insert new vault")
	}

	return vault, nil
}

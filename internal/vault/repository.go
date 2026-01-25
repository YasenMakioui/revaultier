package vault

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/gommon/log"
)

type VaultRepository struct {
	dbconn *sql.DB
}

func NewVaultRepository(db *sql.DB) *VaultRepository {
	return &VaultRepository{dbconn: db}
}

func (r *VaultRepository) GetVault(ctx context.Context, vaultId string, ownerId string) (Vault, error) {

	sqlStmt := "SELECT * FROM vault WHERE id=? and owner_id =?"

	var vault Vault

	err := r.dbconn.QueryRowContext(ctx, sqlStmt, vaultId, ownerId).Scan(&vault.Id, &vault.Owner_id, &vault.Name, &vault.Description, &vault.Created_at)

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

func (r *VaultRepository) InsertVault(ctx context.Context, vault Vault) (Vault, error) {
	sqlStmt := "INSERT INTO vault(id,owner_id,name,description,created_at) values(?,?,?,?,?)"

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, vault.Id, vault.Owner_id, vault.Name, vault.Description, vault.Created_at)

	if err != nil {
		log.Error(err)
		return vault, err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		log.Error("failed inserting data")
		return vault, errors.New("could not insert new vault")
	}

	return vault, nil
}

func (r *VaultRepository) DeleteVault(ctx context.Context, vaultId string, ownerId string) error {
	sqlStmt := "DELETE FROM vault WHERE id = ? and owner_id =?"

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, vaultId)

	if err != nil {
		log.Error(err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		log.Error("tried to delete a non-existent vault")
		return errors.New("vault not found")
	}

	return nil
}

func (r *VaultRepository) UpdateVault(ctx context.Context, name string, description string, vaultId string, owner_id string) error {
	sqlStmt := "UPDATE vault SET name = ?, description = ? WHERE id = ? and owner_id = ?"

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, name, description, vaultId)

	if err != nil {
		log.Error(err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		log.Error("tried to update a non-existent vault")
		return errors.New("vault not found")
	}

	return nil
}

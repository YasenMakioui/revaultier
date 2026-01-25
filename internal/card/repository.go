package card

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/gommon/log"
)

type CardRepository struct {
	dbconn *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{dbconn: db}
}

func (r *CardRepository) GetCard(ctx context.Context, vaultId string, cardId string, ownerId string) (Card, error) {

	sqlStmt := "SELECT C.* FROM card JOIN vault v on c.vault_id = v.id WHERE c.id = ? AND v.id = ? AND v.owner_id = ?"

	var card Card

	err := r.dbconn.QueryRowContext(ctx, sqlStmt, cardId, vaultId, ownerId).Scan(&card.Id, &card.Vault_id, &card.Name, &card.Description, &card.Created_at, &card.File)

	if err != nil {
		return Card{}, err
	}

	return card, err
}

func (r *CardRepository) GetCards(ctx context.Context, vaultId string, ownerId string) ([]Card, error) {

	sqlStmt := "SELECT C.* FROM card JOIN vault v on c.vault_id = ? WHERE v.id = ? AND v.owner_id = ?"

	var cards []Card

	rows, err := r.dbconn.QueryContext(ctx, sqlStmt, vaultId, ownerId)

	for rows.Next() {
		if err != nil {
			break
		}

		var card Card

		err := rows.Scan(&card.Id, &card.Vault_id, &card.Name, &card.Description, &card.Created_at, &card.File)

		if err != nil {
			return cards, nil
		}

		cards = append(cards, card)
	}

	return cards, nil
}

func (r *CardRepository) InsertCard(ctx context.Context, card Card, ownerId string) (Card, error) {

	sqlStmt := "INSERT INTO card(id, vault_id, name, description, created_at, file) SELECT ?, v.id, ?,?,?,? FROM vault v WHERE v.id = ? AND v.owner_id = ?"

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, card.Id, card.Name, card.Description, card.Created_at, card.File, card.Vault_id, ownerId)

	if err != nil {
		log.Error(err)
		return card, err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		log.Error("failed inserting data")
		return card, errors.New("could not insert new vault")
	}

	return card, nil
}

func (r *CardRepository) DeleteCard(ctx context.Context, cardId string, vaultId string, ownerId string) error {
	sqlStmt := "DELETE FROM card WHERE id = ? AND vault_id IN (SELECT v.id FROM vault v WHERE v.id = ? AND v.owner_id = ?)"

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, cardId, vaultId, ownerId)

	if err != nil {
		log.Error(err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		log.Error("tried to delete a non-existent card")
		return errors.New("card not found")
	}

	return nil
}

func (r *CardRepository) UpdateCard(ctx context.Context, name string, description string, file string, vaultId string, ownerId string) error {
	sqlStmt := "UPDATE card SET name = ?, description = ?, file = ? WHERE id = ? AND EXISTS (SELECT 1 FROM vault v WHERE v.id = card.vault_id AND v.id = ?, AND v.owner_id = ?)"

	result, err := r.dbconn.ExecContext(ctx, sqlStmt, name, description, file, vaultId, ownerId)

	if err != nil {
		log.Error(err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		log.Error("failed to update card")
		return errors.New("could not update card")
	}

	return nil
}

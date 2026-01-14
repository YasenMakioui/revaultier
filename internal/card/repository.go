package card

import (
	"database/sql"
)

type CardRepository struct {
	dbconn *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{dbconn: db}
}

//func (r *CardRepository) GetCard(ctx context.Context)

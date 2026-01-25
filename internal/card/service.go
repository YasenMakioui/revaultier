package card

import (
	"context"
	"database/sql"
	"errors"
	"revaultier/configuration"

	"github.com/labstack/gommon/log"
)

type CardService struct {
	cfg            *configuration.Config
	CardRepository *CardRepository
}

func NewCardService(cfg *configuration.Config, cr *CardRepository) *CardService {
	return &CardService{
		cfg:            cfg,
		CardRepository: cr,
	}
}

func (s *CardService) GetCardService(ctx context.Context, vaultId string, cardId string) (Card, error) {

	card, err := s.CardRepository.GetCard(ctx, vaultId, cardId)

	if err != nil {
		if err == sql.ErrNoRows {
			return Card{}, nil
		}
		log.Error("could not get vault")
		return Card{}, errors.New("could not get vault")
	}

	return card, nil
}

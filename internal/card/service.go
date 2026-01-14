package card

import "revaultier/configuration"

type CardService struct {
	cfg            *configuration.Config
	cardRepository *CardRepository
}

func NewCardService(cfg *configuration.Config, cardRepository *CardRepository) *CardService {
	return &CardService{
		cfg:            cfg,
		cardRepository: cardRepository,
	}
}

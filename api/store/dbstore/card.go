package dbstore

import (
	"gorm.io/gorm"
	"itruj/discens-vestigo/api/store"
)

type CardStore struct {
	db *gorm.DB
}

type NewCardStoreParams struct {
	DB *gorm.DB
}

func NewCardStore(params NewCardStoreParams) *CardStore {
	return &CardStore{
		db: params.DB,
	}
}

func (s *CardStore) CreateCard(title string, done bool) error {
	return s.db.Create(&store.Card{
		Title: title,
		Done:  done,
	}).Error
}

func (s *CardStore) GetCard(title string) (*store.Card, error) {
	var card store.Card
	err := s.db.Where("title = ?", title).First(&card).Error

	if err != nil {
		return nil, err
	}

	return &card, err
}

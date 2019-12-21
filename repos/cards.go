package repos

import (
	"github.com/jinzhu/gorm"
	"monster_duel_api/models"
)


//CardRepo struct
type CardRepo struct {
	DBConnection	*gorm.DB
}

// ListCards will get a limited list of card model
func (cr *CardRepo) ListCards(limit int, offset int) []*models.Card {
	var cardModelList []*models.Card
	verifiedLimit := limit

	if limit > 30 {
		verifiedLimit = 30
	}
	
	cr.DBConnection.
		Limit(verifiedLimit).
		Offset(offset).
		Find(&cardModelList)

	return cardModelList
}

// GetCardByName will find card by name
func (cr *CardRepo) GetCardByName(title string) *models.Card {
	var card models.Card
	likeTitle := "%"+title+"%"

	cr.DBConnection.	
		Where("title LIKE = ?", likeTitle).
		First(&card)

	return &card
}

// GetCardByID will find card by ID
func (cr *CardRepo) GetCardByID(id int) *models.Card {
	var card models.Card

	cr.DBConnection.	
		Where("id = ?", id).
		First(&card)

	return &card
}

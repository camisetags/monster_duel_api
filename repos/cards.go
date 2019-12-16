package repos

import (
	"github.com/jinzhu/gorm"
	"monster_duel_api/models"
)


//CardRepo struct
type CardRepo struct {
	DBConnection	*gorm.DB
}

// GetCards will get a limited list of card model
func (cr *CardRepo) GetCards(limit int, offset int) []*models.Card {
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

package repos

import (
	"github.com/jinzhu/gorm"
	"monster_duel_api/models"
)


//CardRepo struct
type CardRepo struct {}

// GetCards will get a limited list of card model
func (CardRepo) GetCards(limit int, dbConnection *gorm.DB) []*models.Card {
	var cardModelList []*models.Card
	dbConnection.Limit(10).Find(&cardModelList)

	return cardModelList
}

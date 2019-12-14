package models

import (
	"github.com/jinzhu/gorm"
	"github.com/camisetags/monster_duel_api_go/schema"
)

// CardModel struct
type CardModel struct {
	gorm.Model
	schema.Card
	CardSets    string `json:"card_sets"`
	CardPrices  string `json:"card_prices"`	
	BanlistInfo string  `json:"banlist_info"`
}
	
// TableName sets table name
func (CardModel) TableName() string {
	return "cards"
}


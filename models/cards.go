package models

import (
	"time"
	
	"github.com/jinzhu/gorm"
	// "monster_duel_api/schema"
)

// CardModel struct
type CardModel struct {
	gorm.Model
	
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Desc        *string    `json:"desc"`
	Atk         *string    `json:"atk"`
	Def         *string    `json:"def"`
	Level       *string    `json:"level"`
	Race        *string    `json:"race"`
	Attribute   *string    `json:"attribute"`

	Archetype   *string    `json:"archetype"`
	Linkval     *string    `json:"linkval"`
	Fname       *string    `json:"fname"`
	Rank        *string    `json:"rank"`
	
	Format      *string    `json:"format"`
	Sort        *string    `json:"sort"`
	La          *string    `json:"la"`
	Scale       *string    `json:"scale"`
	
	Linkmarkers *string  `json:"linkmarkers"`
	CardSets    string `json:"card_sets"`
	CardPrices  string `json:"card_prices"`	
	BanlistInfo string `json:"banlist_info"`
	DeletedAt	*time.Time
}
	
// TableName sets table name
func (CardModel) TableName() string {
	return "cards"
}


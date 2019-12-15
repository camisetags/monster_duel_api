package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// BanList struct
type BanList struct {
	BanTcg *string `json:"ban_tcg"`
	BanOcg *string `json:"ban_ocg"`
}

// CardPrice struct
type CardPrice struct {
	CardmarketPrice *string `json:"cardmarket_price"`
	TcgplayerPrice  *string `json:"tcgplayer_price"`
	EbayPrice       *string `json:"ebay_price"`
	AmazonPrice     *string `json:"amazon_price"`
}

//CardSet struct
type CardSet struct {
	SetName   *string `json:"set_name"`
	SetCode   *string `json:"set_code"`
	SetRarity *string `json:"set_rarity"`
	SetPrice  *string `json:"set_price"`
}

// Card struct
type Card struct {
	gorm.Model
	ID          		int        `json:"id"`
	Name        		string     `json:"name"`
	Type        		string     `json:"type"`
	Desc        		*string    `json:"desc"`
	Atk         		*string    `json:"atk"`
	Def         		*string    `json:"def"`
	Level       		*string    `json:"level"`
	Race        		*string    `json:"race"`
	Attribute   		*string    `json:"attribute"`
	Archetype   		*string    `json:"archetype"`
	Linkval     		*string    `json:"linkval"`
	Fname       		*string    `json:"fname"`
	Rank        		*string    `json:"rank"`
	Format      		*string    `json:"format"`
	Sort        		*string    `json:"sort"`
	La          		*string    `json:"la"`
	Scale       		*string    `json:"scale"`
	
	CardSetsString    	*string    `json:"card_sets" gorm:"column:card_sets"`
	CardPricesString  	*string    `json:"card_prices" gorm:"column:card_prices"`
	LinkmarkersString 	*string    `json:"linkmarkers" gorm:"column:linkmarkers"`
	BanListInfoString 	*string    `json:"banlist_info" gorgm:"banlist_info"`

	DeletedAt	*time.Time
}

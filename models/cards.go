package models

import (
	"fmt"
	"encoding/json"

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

// GetCards will get a limited list of card model
func (CardModel) GetCards(limit int, dbConnection *gorm.DB) []*schema.Card {
	var cardModelList []*CardModel
	dbConnection.Limit(10).Find(&cardModelList)
	convertedCards := convertCardList(cardModelList)
	
	return convertedCards
}

func convertCardList(cards []*CardModel) []*schema.Card {
	var cardList []*schema.Card
	cardsChan := make(chan *schema.Card)

	go func() {
		for _, c := range cards {
			convertedCard, err := parseCardsToModel(c)
			
			if err != nil {
				fmt.Println(err)
			} else {
				cardsChan <- convertedCard
			}
		}
		
		close(cardsChan)
	}()

	for {
		cardConverted, ok := <- cardsChan

		if ok {
			cardList = append(cardList, cardConverted)
		} else {
			break
		}
	}

	return cardList
}

func parseCardsToModel(card *CardModel) (*schema.Card, error) {
	var cardPrices *schema.CardPrice
	var banListInfo *schema.BanList
	var cardSets []*schema.CardSet
	
	err := json.Unmarshal([]byte(card.CardPrices), &cardPrices)
	err = json.Unmarshal([]byte(card.BanlistInfo), &banListInfo)
	err = json.Unmarshal([]byte(card.CardSets), &cardSets)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cardConverted := &schema.Card{
		Name: card.Name,
		Type: card.Type,
		Desc: card.Desc,
		Atk: card.Atk,
		Def: card.Def,
		Level: card.Level,
		Race: card.Race,
		Attribute: card.Attribute,
		Archetype: card.Archetype,
		Linkval: card.Linkval,
		Fname: card.Fname,
		Rank: card.Rank,
		Linkmarkers: card.Linkmarkers,
		Format: card.Format,
		Sort: card.Sort,
		La: card.La,
		Scale: card.Scale,
		CardPrices: cardPrices,
		BanlistInfo: banListInfo,
		CardSets: cardSets,
	}
	
	return cardConverted, nil
}

package repos

import (
	// "fmt"
	// "encoding/json"
	// "sync"

	"github.com/jinzhu/gorm"
	"monster_duel_api/schema"
	// "monster_duel_api/models"
)


//CardRepo struct
type CardRepo struct {}

// GetCards will get a limited list of card model
func (CardRepo) GetCards(limit int, dbConnection *gorm.DB) []*schema.Card {
	var cardModelList []*schema.Card
	dbConnection.Limit(10).Find(&cardModelList)
	// cardList := convertCardList(cardModelList)

	return cardModelList
}

// func convertCardList(cardModels []*models.CardModel) []*schema.Card {
// 	var cardSchemas []*schema.Card
// 	var wg sync.WaitGroup
// 	cardsChan := make(chan *schema.Card)
// 	waitGroupSize := len(cardModels) * 2

// 	wg.Add(waitGroupSize)
	
// 	for _, card := range cardModels {
// 		go func(cardParam *models.CardModel) {
// 			convertedCard, err := parseCardsToModel(*cardParam)
			
// 			if err != nil {
// 				fmt.Printf("Could not convert %s due to %s \n", cardParam.Name, err)
// 				cardsChan <- &schema.Card{ Name: "" }
// 			} else {
// 				cardsChan <- convertedCard
// 			}

// 			wg.Done()
// 		}(card)
// 	}
	
// 	for i := 0; i < len(cardModels); i++ {
// 		go func() {
// 			cardConverted := <- cardsChan

// 			if len(cardConverted.Name) > 0 {
// 				cardSchemas = append(cardSchemas, cardConverted)
// 			}
			
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()	

// 	return cardSchemas
// }

// func parseCardsToModel(card models.CardModel) (*schema.Card, error) {
// 	var cardPrices *schema.CardPrice
// 	var banListInfo *schema.BanList
// 	var cardSets []*schema.CardSet
	
// 	err := json.Unmarshal([]byte(card.CardPrices), &cardPrices)
// 	err = json.Unmarshal([]byte(card.BanlistInfo), &banListInfo)
// 	err = json.Unmarshal([]byte(card.CardSets), &cardSets)

// 	if err != nil {
// 		fmt.Println(card.CardPrices)
// 		fmt.Println(card.BanlistInfo)
// 		fmt.Println(card.CardSets)
// 		fmt.Println(err)
// 	}

// 	cardConverted := &schema.Card{
// 		Name: card.Name,
// 		Type: card.Type,
// 		Desc: card.Desc,
// 		Atk: card.Atk,
// 		Def: card.Def,
// 		Level: card.Level,
// 		Race: card.Race,
// 		Attribute: card.Attribute,
// 		Archetype: card.Archetype,
// 		Linkval: card.Linkval,
// 		Fname: card.Fname,
// 		Rank: card.Rank,
// 		Linkmarkers: card.Linkmarkers,
// 		Format: card.Format,
// 		Sort: card.Sort,
// 		La: card.La,
// 		Scale: card.Scale,
// 		CardPrices: cardPrices,
// 		BanlistInfo: banListInfo,
// 		CardSets: cardSets,
// 	}
	
// 	return cardConverted, nil
// }

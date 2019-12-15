package resolvers

import (
	"context"
	"fmt"
	"encoding/json"

	"monster_duel_api/generated"
	"monster_duel_api/database"
	"monster_duel_api/repos"
	"monster_duel_api/schema"
)

type cardResolver struct{ *Resolver }

// Card resolver
func (r *Resolver) Card() generated.CardResolver {
	return &cardResolver{r}
}

func (cr *cardResolver) CardSets(ctx context.Context, obj *schema.Card) ([]*schema.CardSet, error) {
	var cs []*schema.CardSet
	cardSetsBlob := []byte(*obj.CardSetsString)

	err := json.Unmarshal(cardSetsBlob, &cs)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return cs, nil
}

func (cr *cardResolver) CardPrices(ctx context.Context, obj *schema.Card) (*schema.CardPrice, error) {
	var cardPrice *schema.CardPrice
	cardPricesBlob := []byte(*obj.CardPricesString)

	err := json.Unmarshal(cardPricesBlob, &cardPrice)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return cardPrice, nil
}

// Linkmarkers resolver
func (cr *cardResolver) Linkmarkers(ctx context.Context, obj *schema.Card) ([]*string, error) {
	var linkmarker []*string
	linkMarkerBlob := []byte(*obj.LinkmarkersString)

	err := json.Unmarshal(linkMarkerBlob, &linkmarker)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return linkmarker, nil
}

// BanlistInfo resolver
func (cr *cardResolver) BanlistInfo(ctx context.Context, obj *schema.Card) (*schema.BanList, error) {
	var banList *schema.BanList
	banListBlob := []byte(*obj.BanListInfoString)

	err := json.Unmarshal(banListBlob, &banList)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return banList, nil
}

// Cards resolver to list cards
func (r *queryResolver) Cards(ctx context.Context) ([]*schema.Card, error) {
	db := database.GetDBConnection()
	repo := &repos.CardRepo{}
	cards := repo.GetCards(10, db)

	return cards, nil
}

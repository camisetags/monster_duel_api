package resolvers

import (
	"context"
	"fmt"
	"encoding/json"

	"monster_duel_api/generated"
	"monster_duel_api/database"
	"monster_duel_api/repos"
	"monster_duel_api/models"
)

type cardResolver struct{ *Resolver }

// Card resolver
func (r *Resolver) Card() generated.CardResolver {
	return &cardResolver{r}
}

func (cr *cardResolver) ID(ctx context.Context, obj *models.Card) (int, error) {
	return int(obj.Model.ID), nil
}

func (cr *cardResolver) CardSets(ctx context.Context, obj *models.Card) ([]*models.CardSet, error) {
	if obj.CardSetsString == nil {
		return nil, nil
	}

	var cs []*models.CardSet
	cardSetsBlob := []byte(*obj.CardSetsString)
	err := json.Unmarshal(cardSetsBlob, &cs)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return cs, nil
}

func (cr *cardResolver) CardPrices(ctx context.Context, obj *models.Card) (*models.CardPrice, error) {
	if obj.CardPricesString == nil {
		return nil, nil
	}

	var cardPrice *models.CardPrice
	cardPricesBlob := []byte(*obj.CardPricesString)
	err := json.Unmarshal(cardPricesBlob, &cardPrice)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return cardPrice, nil
}

// Linkmarkers resolver
func (cr *cardResolver) Linkmarkers(ctx context.Context, obj *models.Card) ([]*string, error) {
	if obj.LinkmarkersString == nil {
		return nil, nil
	}

	var linkmarker []*string
	linkMarkerBlob := []byte(*obj.LinkmarkersString)
	err := json.Unmarshal(linkMarkerBlob, &linkmarker)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, err
	}

	return linkmarker, nil
}

// BanlistInfo resolver
func (cr *cardResolver) BanlistInfo(ctx context.Context, obj *models.Card) (*models.BanList, error) {
	if obj.BanListInfoString == nil {
		return nil, nil
	}

	var banList *models.BanList
	banListBlob := []byte(*obj.BanListInfoString)
	err := json.Unmarshal(banListBlob, &banList)

	if err != nil {
		fmt.Println("Could not convert card_sets")
		return nil, nil
	}

	return banList, nil
}

// Cards resolver to list cards
func (r *queryResolver) Cards(ctx context.Context, limit *int, offset *int) ([]*models.Card, error) {
	db := database.GetDBConnection()
	repo := &repos.CardRepo{ DBConnection: db }

	verifiedLimit := setDefaultIfNil(limit, 30).(int)
	verifiedOffset := setDefaultIfNil(offset, 0).(int)

	cards := repo.GetCards(verifiedLimit, verifiedOffset)

	return cards, nil
}

func (r *queryResolver) Card(ctx context.Context, name *string, id *int) (*models.Card, error) {
	panic("not implemented")
}

func setDefaultIfNil(val *int, defaultValue int) interface{} {
	if val == nil {
		return defaultValue
	}
	return *val
}

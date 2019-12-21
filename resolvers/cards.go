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
		return nil, err
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
		fmt.Println("Could not convert card_prices")
		return nil, err
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
		fmt.Println("Could not convert linkmarkers")
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
		fmt.Println("Could not convert banlist_info")
		return nil, err
	}

	return banList, nil
}

// Cards resolver to list cards
func (r *queryResolver) Cards(ctx context.Context, limit *int, offset *int) ([]*models.Card, error) {
	repo := &repos.CardRepo{ 
		DBConnection: database.GetDBConnection(), 
	}

	verifiedLimit := setDefaultIfNil(limit, 30).(int)
	verifiedOffset := setDefaultIfNil(offset, 0).(int)

	cards := repo.ListCards(verifiedLimit, verifiedOffset)

	return cards, nil
}

func (r *queryResolver) CardsByName(ctx context.Context, name string) ([]*models.Card, error) {
	panic("not implemented")
}

func (r *queryResolver) Card(ctx context.Context, id int) (*models.Card, error) {
	repo := &repos.CardRepo{ 
		DBConnection: database.GetDBConnection(),
	}

	return repo.GetCardByID(id), nil

	// return nil, errors.New("Id must be set")
}

func setDefaultIfNil(val *int, defaultValue int) interface{} {
	if val == nil {
		return defaultValue
	}
	return *val
}

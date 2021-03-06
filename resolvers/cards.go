package resolvers

import (
	"context"
	"fmt"
	"encoding/json"
	"errors"

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
func (r *queryResolver) Cards(ctx context.Context, input *models.CardSearchInput) ([]*models.Card, error) {
	repo := &repos.CardRepo{ 
		DBConnection: database.GetDBConnection(), 
	}

	verifiedLimit, verifiedOffset := setPaginationLimits(input.Limit, input.Offset)
	var cards []*models.Card

	if input.By != nil && input.By.Name != nil {
		cards = repo.ListCardByName(
			*input.By.Name, 
			verifiedLimit,
			verifiedOffset,
		)
	} else {
		cards = repo.ListCards(verifiedLimit, verifiedOffset)
	}

	return cards, nil
}


func (r *queryResolver) Card(ctx context.Context, id *int, name *string) (*models.Card, error) {
	repo := &repos.CardRepo{ 
		DBConnection: database.GetDBConnection(),
	}

	if id != nil {
		return repo.GetCardByID(*id), nil
	}

	if name != nil {
		return repo.GetCardByName(*name), nil
	}

	return nil, errors.New("Id or name must be set")
}

func setPaginationLimits(limit *int, offset *int) (int, int) {
	verifiedLimit := setDefaultIfNil(limit, 30).(int)
	verifiedOffset := setDefaultIfNil(offset, 0).(int)

	return verifiedLimit, verifiedOffset
}

func setDefaultIfNil(val *int, defaultValue int) interface{} {
	if val == nil {
		return defaultValue
	}
	return *val
}

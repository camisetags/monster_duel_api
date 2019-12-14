package resolvers

import (
	"context"

	"github.com/camisetags/monster_duel_api_go/database"
	"github.com/camisetags/monster_duel_api_go/models"
	"github.com/camisetags/monster_duel_api_go/schema"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.


func (r *queryResolver) Cards(ctx context.Context) ([]*schema.Card, error) {
	db := database.GetDBConnection()
	cm := &models.CardModel{}
	cards := cm.GetCards(10, db)	

	

	return , nil
}

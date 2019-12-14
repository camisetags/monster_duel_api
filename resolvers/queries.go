package resolvers

import (
	"github.com/camisetags/monster_duel_api_go/generated"
)

// Resolver struct
type Resolver struct{}

type queryResolver struct{ *Resolver }

// Query func
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

package resolvers

import (
	"monster_duel_api/generated"
)

// Resolver struct
type Resolver struct{}

type queryResolver struct{ *Resolver }

// Query func
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

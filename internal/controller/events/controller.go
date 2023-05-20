package events

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
)

type repository interface {
	GetBySport(ctx context.Context, sport string, limit int, page int) *model.PaginationResponse
	GetByQuery(ctx context.Context, query string, limit, page int) *model.PaginationResponse
	GetByTournament(ctx context.Context, tournament string, limit, page int) *model.PaginationResponse
	GetByID(ctx context.Context, id string) (*model.Event, error)
}

type Controller struct {
	r repository
}

func New(r repository) *Controller {
	return &Controller{r}
}

func (c *Controller) GetBySport(ctx context.Context, sport string, limit int, page int) *model.PaginationResponse {
	return c.r.GetBySport(ctx, sport, limit, page)
}

func (c *Controller) GetByQuery(ctx context.Context, query string, limit, page int) *model.PaginationResponse {
	return c.r.GetByQuery(ctx, query, limit, page)
}

func (c *Controller) GetByTournament(ctx context.Context, tournament string, limit, page int) *model.PaginationResponse {
	return c.r.GetByTournament(ctx, tournament, limit, page)
}

func (c *Controller) GetByID(ctx context.Context, id string) (*model.Event, error) {
	return c.r.GetByID(ctx, id)
}

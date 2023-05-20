package tournaments

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
)

type repository interface {
	Get(ctx context.Context) *[]model.Tournament
}

type Controller struct {
	r repository
}

func New(r repository) *Controller {
	return &Controller{r}
}

func (c *Controller) Get(ctx context.Context) *[]model.Tournament {
	return c.r.Get(ctx)
}

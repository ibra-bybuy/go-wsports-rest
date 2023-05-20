package banners

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
)

type repository interface {
	Get(ctx context.Context) *model.BannerList
}

type Controller struct {
	r repository
}

func New(r repository) *Controller {
	return &Controller{r}
}

func (c *Controller) Get(ctx context.Context) *model.BannerList {
	return c.r.Get(ctx)
}

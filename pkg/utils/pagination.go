package utils

import (
	"math"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
)

func BuildPagination(totalItems int64, page, limit int) *model.Pagination {
	return &model.Pagination{
		Page:       page,
		Limit:      limit,
		TotalItems: int(totalItems),
		TotalPages: int(math.Ceil(float64(totalItems) / float64(limit))),
	}
}

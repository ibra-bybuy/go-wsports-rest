package banners

import (
	"context"
	"fmt"

	"github.com/ibra-bybuy/go-wsports-events/pkg/constants"
	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"github.com/ibra-bybuy/go-wsports-events/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) Get(ctx context.Context) *model.BannerList {
	response := model.BannerList{}

	collection, ctx, cancel := m.getCollection()
	defer cancel()

	/*
	* Retrieving unique avatar
	* That are mma and not equal to ""
	 */
	findFilter := bson.D{{Key: "startAt", Value: bson.D{{Key: "$gt", Value: utils.StartDate("")}}}, {Key: "sport", Value: constants.MMA_TYPE}, {Key: "avatarUrl", Value: bson.D{{Key: "$ne", Value: ""}}}}
	values, err := collection.Distinct(ctx, "avatarUrl", findFilter)

	if err != nil {
		return &response
	}
	avatars := []string{}
	for _, value := range values {
		avatars = append(avatars, fmt.Sprintf("%s", value))
	}

	uniqueAvatars := utils.StrListUnique(avatars)
	if len(uniqueAvatars) < 1 {
		return &response
	}

	/*
	* Retrieving events by the avatar
	 */

	findFilter = bson.D{{Key: "avatarUrl", Value: bson.D{{Key: "$in", Value: uniqueAvatars}}}}
	cursor, err := collection.Find(ctx, findFilter)

	if err != nil {
		return &response
	}

	events := []model.Event{}
	cursor.All(ctx, &events)
	for _, event := range events {
		contains := response.Contains(event.AvatarURL)
		if contains {
			continue
		}

		response = append(response, model.Banner{
			Code:      event.Name,
			Name:      event.Name,
			AvatarURL: event.AvatarURL,
		})
	}

	return &response
}

package model

type Banner struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

type BannerList []Banner

func (bl *BannerList) Contains(avatar string) bool {
	for _, b := range *bl {
		if b.AvatarURL == avatar {
			return true
		}
	}

	return false
}

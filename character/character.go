package character

import (
	"github.com/dywoq/genshinapi/statistics/element"
	"github.com/dywoq/genshinapi/statistics/model"
	"github.com/dywoq/genshinapi/statistics/quality"
	"github.com/dywoq/genshinapi/statistics/region"
	"github.com/dywoq/genshinapi/statistics/version"
	"github.com/dywoq/genshinapi/statistics/weapon"
	"github.com/dywoq/genshinapi/statistics/wish"
)

type Character struct {
	Name    string          `json:"name"`
	Quality quality.Type    `json:"quality"`
	Element element.Type    `json:"element"`
	Weapon  weapon.Type     `json:"weapon"`
	Region  region.Type     `json:"region"`
	Model   model.Type      `json:"model"`
	Version version.Version `json:"version"`
	Wish    wish.Type       `json:"wish"`
}

func Correct(ch *Character) bool {
	checks := getChecks(ch)
	for check := range checks {
		err := check.Check()
		if err != nil {
			return false
		}
	}
	return true
}

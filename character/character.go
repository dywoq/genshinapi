package character

import (
	"github.com/dywoq/genshinapi/statistics/element"
	"github.com/dywoq/genshinapi/statistics/kind"
	"github.com/dywoq/genshinapi/statistics/model"
	"github.com/dywoq/genshinapi/statistics/quality"
	"github.com/dywoq/genshinapi/statistics/region"
	"github.com/dywoq/genshinapi/statistics/version"
	"github.com/dywoq/genshinapi/statistics/weapon"
)

type Character struct {
	Quality quality.Type    `json:"quality"`
	Element element.Type    `json:"element"`
	Weapon  weapon.Type     `json:"weapon"`
	Region  region.Type     `json:"region"`
	Model   model.Type      `json:"model"`
	Version version.Version `json:"version"`
	Kind    kind.Type       `json:"kind"`
}

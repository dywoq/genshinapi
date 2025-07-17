package character

import (
	"github.com/dywoq/genshinapi/statistics"
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

func New(q quality.Type, e element.Type, w weapon.Type, r region.Type, m model.Type, v version.Version, k kind.Type) (*Character, error) {
	var strChecks = map[statistics.Checker[string]]struct{}{
		q: {},
		e: {},
		w: {},
		r: {},
		m: {},
		k: {},
	}

	for strCheck := range strChecks {
		err := strCheck.Check()
		if err != nil {
			return &Character{}, err
		}
	}

	if err := v.Check(); err != nil {
		return &Character{}, err
	}

	return &Character{q, e, w, r, m, v, k}, nil
}

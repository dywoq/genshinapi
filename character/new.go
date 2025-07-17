package character

import (
	"github.com/dywoq/genshinapi/statistics"
	"github.com/dywoq/genshinapi/statistics/element"
	"github.com/dywoq/genshinapi/statistics/model"
	"github.com/dywoq/genshinapi/statistics/quality"
	"github.com/dywoq/genshinapi/statistics/region"
	"github.com/dywoq/genshinapi/statistics/version"
	"github.com/dywoq/genshinapi/statistics/weapon"
	"github.com/dywoq/genshinapi/statistics/wish"
)

func New(n string, q quality.Type, e element.Type, w weapon.Type, r region.Type, m model.Type, v version.Version, wish wish.Type) (*Character, error) {
	checks := map[statistics.Checker]struct{}{
		q: {},
		e: {},
		w: {},
		r: {},
		m: {},
		wish: {},
		v: {},
	}

	for check := range checks {
		err := check.Check()
		if err != nil {
			return &Character{}, err
		}
	}
	return &Character{n, q, e, w, r, m, v, wish}, nil
}

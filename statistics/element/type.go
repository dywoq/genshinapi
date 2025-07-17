package element

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	validTypes := map[Type]struct{}{
		Anemo:   {},
		Cryo:    {},
		Dendro:  {},
		Electro: {},
		Geo:     {},
		Hydro:   {},
		Pyro:    {},
		None:    {},
	}

	if _, ok := validTypes[t]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

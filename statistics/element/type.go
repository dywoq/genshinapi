package element

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	switch (t) {
	case Anemo, Cryo, Dendro, Electro, Geo, Hydro, Pyro, None:
		return nil
	default:
		return statistics.ErrCheckFailed
	}
}

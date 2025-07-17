package weapon

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	var validTypes = map[Type]struct{}{
		Sword:    {},
		Bow:      {},
		Catalyst: {},
		Polearm:  {},
		Claymore: {},
	}

	if _, ok := validTypes[t]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

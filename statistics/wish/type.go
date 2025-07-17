package wish

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	validTypes := map[Type]struct{}{
		EventExclusive:                                 {},
		CannotBeObtainedFromWishes:                     {},
		CannotBeObtainedFromCharacterWeaponEventWishes: {},
	}

	if _, ok := validTypes[t]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

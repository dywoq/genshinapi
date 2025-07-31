package wish

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	switch (t) {
	case EventExclusive, Standard, CannotBeObtainedFromWishes, CannotBeObtainedFromCharacterWeaponEventWishes:
		return nil
	default:
		return statistics.ErrCheckFailed
	}
}

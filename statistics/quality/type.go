package quality

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	var validTypes = map[Type]struct{}{
		FiveStar:  {},
		FourStar:  {},
		ThreeStar: {},
		TwoStar:   {},
		OneStar:   {},
	}

	if _, ok := validTypes[t]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

package model

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	var validTypes = map[Type]struct{}{
		TallMale:     {},
		TallFemale:   {},
		MediumMale:   {},
		MediumFemale: {},
		ShortFemale:  {},
	}

	if _, ok := validTypes[t]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

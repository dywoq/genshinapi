package model

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	switch (t) {
	case TallMale, TallFemale, MediumMale, MediumFemale, ShortFemale:
		return nil
	default:
		return statistics.ErrCheckFailed
	}
}

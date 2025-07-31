package quality

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	switch (t) {
	case FiveStar, FourStar, ThreeStar, TwoStar, OneStar:
		return nil
	default:
		return statistics.ErrCheckFailed
	}
}

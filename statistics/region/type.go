package region

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	switch (t) {
	case Fountaine, Inazuma, Liyue, Mondstadt, Natlan, Snezhnaya, Sumeru:
		return nil
	default:
		return statistics.ErrCheckFailed
	}
}

package region

import "github.com/dywoq/genshinapi/statistics"

type Type string

func (t Type) Check() error {
	validTypes := map[Type]struct{}{
		Fountaine: {},
		Inazuma:   {},
		Liyue:     {},
		Mondstadt: {},
		Natlan:    {},
		Snezhnaya: {},
		Sumeru:    {},
	}

	if _, ok := validTypes[t]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

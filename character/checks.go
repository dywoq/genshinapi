package character

import "github.com/dywoq/genshinapi/statistics"

func getChecks(ch *Character) map[statistics.Checker]struct{} {
	return map[statistics.Checker]struct{}{
		ch.Quality: {},
		ch.Element: {},
		ch.Weapon:  {},
		ch.Region:  {},
		ch.Model:   {},
		ch.Wish:    {},
		ch.Version: {},
	}
}

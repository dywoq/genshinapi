package version

import "github.com/dywoq/genshinapi/statistics"

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
}

func (v Version) Check() error {
	validTypes := map[Version]struct{}{
		Version10: {},
		Version11: {},
		Version12: {},
		Version13: {},
		Version14: {},
		Version15: {},
		Version16: {},

		Version20: {},
		Version21: {},
		Version22: {},
		Version23: {},
		Version24: {},
		Version25: {},
		Version26: {},
		Version27: {},
		Version28: {},

		Version30: {},
		Version31: {},
		Version32: {},
		Version33: {},
		Version34: {},
		Version35: {},
		Version36: {},
		Version37: {},
		Version38: {},

		Version40: {},
		Version41: {},
		Version42: {},
		Version43: {},
		Version44: {},
		Version45: {},
		Version46: {},
		Version47: {},
		Version48: {},

		Version50: {},
		Version51: {},
		Version52: {},
		Version53: {},
		Version54: {},
		Version55: {},
		Version56: {},
		Version57: {},
	}

	if _, ok := validTypes[v]; ok {
		return nil
	}
	return statistics.ErrCheckFailed
}

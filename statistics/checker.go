package statistics

import "errors"

var ErrCheckFailed = errors.New("github.com/dywoq/genshinapi/statistics: check failed")

type Checker interface {
	Check() error
}

package flag

import (
	"github.com/spf13/pflag"
)

type optionFn func(*flag)

func WithFlagSets(set *pflag.FlagSet) optionFn {
	return func(f *flag) {
		if f.sets == nil {
			f.sets = set
		}
	}
}

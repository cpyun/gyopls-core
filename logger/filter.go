package logger

import (
	"github.com/cpyun/gyopls-core/logger/level"
)

const fuzzyStr = "***"

// filterOptionFunc is filter option.
type filterOptionFunc func(*Logger)

// FilterWithLevel with filter level.
func FilterWithLevel(lvl level.Level) filterOptionFunc {
	return func(o *Logger) {
		f := func(level level.Level, _ ...any) bool {
			return lvl < o.opts.level
		}

		o.filter = append(o.filter, f)
	}
}

// FilterWithKey with filter key.
func FilterWithKey(key ...string) filterOptionFunc {
	return func(o *Logger) {
		var keyMap = make(map[string]bool, len(key))
		for _, v := range key {
			keyMap[v] = true
		}

		f := func(_ level.Level, keyVals ...any) bool {
			for i := 0; i < len(keyVals); {
				if i == len(keyVals)-1 || i+1 >= len(keyVals) {
					return false
				}

				key, okKey := keyVals[i].(string)
				if !okKey {
					i++
					continue
				}

				if _, okMap := keyMap[key]; okMap {
					keyVals[i+1] = fuzzyStr
				}
				i += 2
			}
			return true
		}

		o.filter = append(o.filter, f)
	}
}

// FilterWithValue with filter value.
func FilterWithValue(values ...string) filterOptionFunc {
	return func(o *Logger) {
		var keyMap = make(map[string]bool, len(values))
		for _, v := range values {
			if v == fuzzyStr {
				continue
			}
			keyMap[v] = true
		}

		f := func(_ level.Level, keyVals ...any) bool {
			for i := 0; i < len(keyVals); {
				if i == len(keyVals)-1 || i+1 >= len(keyVals) {
					return false
				}

				if _, okKey := keyVals[i].(string); !okKey {
					i++
					continue
				}

				// 判断value是否为string类型
				val, okVal := keyVals[i+1].(string)
				if !okVal {
					i++
					continue
				}

				if _, okMap := keyMap[val]; okMap {
					keyVals[i+1] = fuzzyStr
				}
				i += 2
			}
			return true
		}

		o.filter = append(o.filter, f)
	}
}

// FilterWithFunc with filter func.
func FilterWithFunc(f func(_ level.Level, keyvals ...any) bool) filterOptionFunc {
	return func(o *Logger) {
		o.filter = append(o.filter, f)
	}
}

// Filter filter.
func (l *Logger) Filter(fs ...filterOptionFunc) *Logger {
	for _, f := range fs {
		f(l)
	}
	return l
}

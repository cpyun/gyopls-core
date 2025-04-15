package logger

const fuzzyStr = "***"

// filterOptionFunc is filter option.
type filterOptionFunc func(*Logger)

type filterOptions struct {
	attrs []func(keyVals ...any) bool
}

// FilterWithKey with filter key.
func FilterWithKey(key ...string) filterOptionFunc {
	fn := func(keyVals ...any) bool {
		var keyMap = make(map[string]bool, len(key))
		for _, v := range key {
			keyMap[v] = true
		}
		for i := 0; i < len(keyVals); {
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

	return FilterWithFunc(fn)
}

// FilterWithValue with filter value.
func FilterWithValue(values ...string) filterOptionFunc {
	fn := func(keyVals ...any) bool {
		var keyMap = make(map[string]bool, len(values))
		for _, v := range values {
			if v == fuzzyStr {
				continue
			}
			keyMap[v] = true
		}
		for i := 0; i < len(keyVals); {
			if i+1 >= len(keyVals) {
				return false
			}

			if _, okKey := keyVals[i].(string); !okKey {
				i++
				continue
			}

			// 判断value是否为string类型
			val, okVal := keyVals[i+1].(string)
			if !okVal {
				i += 2
				continue
			}

			if _, okMap := keyMap[val]; okMap {
				keyVals[i+1] = fuzzyStr
			}
			i += 2
		}
		return true
	}

	return FilterWithFunc(fn)
}

// FilterWithFunc with filter func.
func FilterWithFunc(f func(keyvals ...any) bool) filterOptionFunc {
	return func(o *Logger) {
		o.filter.attrs = append(o.filter.attrs, f)
	}
}

// checkFilter 校验过滤
func (t *Logger) checkFilter(args ...any) []any {
	for _, f := range t.filter.attrs {
		if f != nil && !f(args...) {
			break
		}
	}
	return args
}

// Filter filter.
func (l *Logger) Filter(fs ...filterOptionFunc) *Logger {
	for _, f := range fs {
		f(l)
	}
	return l
}

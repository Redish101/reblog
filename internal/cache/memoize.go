package cache

import "fmt"

func Memoize[T any, R any](f func(T) R) func(T) R {
	cache := make(map[string]R)

	return func(arg T) R {
		key := fmt.Sprintf("%v", arg)

		if val, found := cache[key]; found {
			return val
		}

		result := f(arg)

		cache[key] = result

		return result
	}
}

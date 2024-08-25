package cache

import "sync"

func Memoize[T comparable, R any](f func(T) R) func(T) R {
	cache := make(map[T]R)
	var mu sync.Mutex

	return func(x T) R {
		mu.Lock()
		if result, found := cache[x]; found {
			mu.Unlock()
			return result
		}
		mu.Unlock()

		result := f(x)

		mu.Lock()
		cache[x] = result
		mu.Unlock()

		return result
	}
}

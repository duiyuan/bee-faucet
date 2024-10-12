package utils

func GetKeys[K comparable, V any](obj map[K]V) []K {
	ret := make([]K, 0, len(obj))
	for k := range obj {
		ret = append(ret, k)
	}
	return ret
}

func GetValues[K comparable, V any](obj map[K]V) []V {
	ret := make([]V, 0, len(obj))
	for _, v := range obj {
		ret = append(ret, v)
	}
	return ret
}

func Reduce[T any, U any](items []T, initial U, reducer func(U, T) U) U {
	accumulated := initial
	for _, item := range items {
		accumulated = reducer(accumulated, item)
	}
	return accumulated
}

func Map[T any, U any](items []T, mapper func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

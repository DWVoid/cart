package algo

func Select[T any](source []T, apply func(T) bool) (result []T) {
	for _, v := range source {
		if apply(v) {
			result = append(result, v)
		}
	}
	return
}

func Map[T any, U any](source []T, apply func(T) U) (result []U) {
	result = make([]U, len(source))
	for i, v := range source {
		result[i] = apply(v)
	}
	return
}

func TryMap[T any, U any](source []T, apply func(T) (U, error)) (result []U, err error) {
	result = make([]U, len(source))
	for i, v := range source {
		result[i], err = apply(v)
		if err != nil {
			result = nil
			return
		}
	}
	return
}

func MapMap[TK, UK comparable, TV, UV any](source map[TK]TV, apply func(TK, TV) (UK, UV)) (result map[UK]UV) {
	result = make(map[UK]UV, len(source))
	for k, v := range source {
		a, b := apply(k, v)
		result[a] = b
	}
	return
}

func TryMapMap[TK, UK comparable, TV, UV any](source map[TK]TV, apply func(TK, TV) (UK, UV, error)) (map[UK]UV, error) {
	result := make(map[UK]UV, len(source))
	for k, v := range source {
		a, b, err := apply(k, v)
		if err != nil {
			return nil, err
		}
		result[a] = b
	}
	return result, nil
}

func Reduce[T any, U any](source []T, init U, apply func(U, T) U) (result U) {
	result = init
	for _, v := range source {
		result = apply(result, v)
	}
	return
}

func TryReduce[T any, U any](source []T, init U, apply func(U, T) (U, error)) (result U, err error) {
	result = init
	for _, v := range source {
		result, err = apply(result, v)
		if err != nil {
			return
		}
	}
	return
}

func Flatten[T any](source [][]T) (result []T) {
	for _, v := range source {
		result = append(result, v...)
	}
	return
}

func GroupBy[T any, U comparable](source []T, key func(T) U) (result map[U][]T) {
	result = make(map[U][]T)
	for _, v := range source {
		k := key(v)
		result[k] = append(result[k], v)
	}
	return
}

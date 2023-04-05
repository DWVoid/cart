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

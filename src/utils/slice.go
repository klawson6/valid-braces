package utils

func Contains[T comparable](s []T, e T) bool {
	for _, t := range s {
		if t == e {
			return true
		}
	}
	return false
}

package mapx

// Keys returns a slice of all keys in m.
func Keys[M interface{ ~map[K]V }, K comparable, V any](m M) []K {
	if m == nil {
		return nil
	}

	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Values returns a slice of all values in m.
func Values[M interface{ ~map[K]V }, K comparable, V any](m M) []V {
	if m == nil {
		return nil
	}

	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

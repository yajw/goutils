package val

// nilOr if ptr is nil, return defaultValue
func NilOr[T any, PT interface{ *T }](ptr PT, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

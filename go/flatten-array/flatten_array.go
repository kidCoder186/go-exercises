package flatten

// Flatten returns a single flattened list from input.
func Flatten(input interface{}) []interface{} {
	res := []interface{}{}
	switch val := input.(type) {
	case []interface{}:
		for _, v := range val {
			res = append(res, Flatten(v)...)
		}
	default:
		if val != nil {
			res = append(res, val)
		}
	}
	return res
}

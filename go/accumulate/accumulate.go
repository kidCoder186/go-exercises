package accumulate

// Accumulate uses `converter` to transmit input.
func Accumulate(input []string, converter func(string) string) []string {
	res := []string{}
	for _, v := range input {
		res = append(res, converter(v))
	}
	return res
}

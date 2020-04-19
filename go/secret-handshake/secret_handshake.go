package secret

var (
	decode = []string{"wink", "double blink", "close your eyes", "jump"}
)

// Handshake decodes `code` input to list of corresponding string.
func Handshake(code uint) []string {
	res := []string{}
	for i := 0; i < 4; i++ {
		if (code & (1 << i)) == (1 << i) {
			res = append(res, decode[i])
		}
	}
	if (code & 16) == 16 {
		for i, j := 0, len(res)-1; i < len(res)/2; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}

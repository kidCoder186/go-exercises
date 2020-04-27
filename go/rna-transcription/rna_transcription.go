package strand

var (
	DNA2RNA = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
)

// ToRNA transforms a DNA strand to its correspongdin RNA.
func ToRNA(dna string) string {
	res := []rune(dna)
	for i, v := range dna {
		res[i] = DNA2RNA[v]
	}
	return string(res)
}

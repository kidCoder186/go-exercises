package scale

import "strings"

var (
	chromatics = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flats      = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

func gen(t string, interval string, collection []string) []string {
	res := []string{}
	tonic := []rune(t)
	if tonic[0] >= 'a' && tonic[0] <= 'g' {
		tonic[0] = tonic[0] - 32
	}
	pos := 0
	for i := 0; i < 12; i++ {
		if strings.Compare(string(tonic), collection[i]) == 0 {
			pos = i
			break
		}
	}
	res = append(res, string(tonic))
	if strings.Compare(interval, "") == 0 {
		interval = "mmmmmmmmmmmm"
	}
	for _, v := range interval {
		switch v {
		case 'M':
			pos += 2
		case 'm':
			pos++
		case 'A':
			pos += 3
		}
		res = append(res, collection[pos%12])
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return res
}

// Scale returns
func Scale(tonic string, interval string) []string {
	res := []string{}
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		res = gen(tonic, interval, chromatics)
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		res = gen(tonic, interval, flats)
	}
	return res
}

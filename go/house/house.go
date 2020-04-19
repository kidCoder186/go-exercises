package house

import "strings"

var (
	Nouns = []string{
		"the house that Jack built",
		"the malt",
		"the rat",
		"the cat",
		"the dog",
		"the cow with the crumpled horn",
		"the maiden all forlorn",
		"the man all tattered and torn",
		"the priest all shaven and shorn",
		"the rooster that crowed in the morn",
		"the farmer sowing his corn",
		"the horse and the hound and the horn",
	}
	Actions = []string{
		"",
		"that lay in ",
		"that ate ",
		"that killed ",
		"that worried ",
		"that tossed ",
		"that milked ",
		"that kissed ",
		"that married ",
		"that woke ",
		"that kept ",
		"that belonged to ",
	}
)

// Verse returns the k-th verse of the song.
func Verse(k int) string {
	res := []string{"This is "}
	for i := k - 1; i >= 0; i-- {
		res = append(res, Nouns[i])
		res = append(res, "\n")
		res = append(res, Actions[i])
	}
	res[len(res)-2] = "."
	res = res[:len(res)-1]
	return strings.Join(res, "")
}

// Song returns all lyrics of the song.
func Song() string {
	res := []string{}
	for i := 1; i <= 12; i++ {
		res = append(res, Verse(i))
		res = append(res, "\n\n")
	}
	res = res[:len(res)-1]
	return strings.Join(res, "")
}

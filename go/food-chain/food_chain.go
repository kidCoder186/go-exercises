package foodchain

import (
	"strings"
)

var (
	phrases = [][3]string{
		{"\n", "\n"},
		{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", ""},
		{"spider", "It wriggled and jiggled and tickled inside her.", "to catch the fly."},
		{"bird", "How absurd to swallow a bird!", "to catch the spider that wriggled and jiggled and tickled inside her."},
		{"cat", "Imagine that, to swallow a cat!", "to catch the bird."},
		{"dog", "What a hog, to swallow a dog!", "to catch the cat."},
		{"goat", "Just opened her throat and swallowed a goat!", "to catch the dog."},
		{"cow", "I don't know how she swallowed a cow!", "to catch the goat."},
		{"horse", "She's dead, of course!", ""},
	}
	storageVerse [9]string
)

// Verse returns the lyrics of id-th verse of the song.
func Verse(id int) string {
	if storageVerse[id] != "" {
		return storageVerse[id]
	}
	if id == 1 || id == 8 {
		storageVerse[id] = "I know an old lady who swallowed a " + phrases[id][0] + ".\n" + phrases[id][1]
		return storageVerse[id]
	}
	var res []string
	sentences := strings.Split(Verse(id-1), "\n")
	res = append(res, sentences[0][:35]+phrases[id][0]+".")
	res = append(res, phrases[id][1])
	res = append(res, "She swallowed the "+phrases[id][0]+" "+phrases[id][2])
	if id == 2 {
		res = append(res, sentences[1:]...)
	} else {
		res = append(res, sentences[2:]...)
	}
	storageVerse[id] = strings.Join(res, "\n")
	return storageVerse[id]
}

// Verses returns the lyrics of the song from `from-th` verse to `to-th` verse.
func Verses(from, to int) string {
	var res []string
	for i := from; i <= to; i++ {
		res = append(res, Verse(i))
	}
	return strings.Join(res, "\n\n")
}

// Song returns the lyrics of the song 'I Know an Old Lady Who Swallowed a Fly'.
func Song() string {
	return Verses(1, len(phrases)-1)
}

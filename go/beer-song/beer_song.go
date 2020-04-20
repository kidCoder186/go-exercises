package beer

import (
	"fmt"
)

// Verse returns the lyrics of id-th verse of the song.
func Verse(id int) (string, error) {
	if id > 2 && id <= 99 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", id, id, id-1), nil
	} else if id == 2 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", id, id, id-1), nil

	} else if id == 1 {
		return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", id, id), nil
	} else if id == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	} else {
		return "", fmt.Errorf("id out of range")
	}
}

// Verses returns the lyrics of the song from `from-th` verse to `to-th` verse.
func Verses(from, to int) (string, error) {
	if from < to {
		return "", fmt.Errorf("from verse < to verse")
	}
	var res string
	for i := from; i >= to; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}
		res += v + "\n"
	}
	return res, nil
}

// Song returns the lyrics of the song '99 Bottles of Beer on the Wall'.
func Song() string {
	res, err := Verses(99, 0)
	if err != nil {
		return ""
	}
	return res
}

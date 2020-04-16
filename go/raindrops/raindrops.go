package raindrops

import "strconv"

// Convert changes a number into a string that contains raindrop sounds
func Convert(num int) string {
	var res string
	if num%3 == 0 {
		res = res + "Pling"
	}
	if num%5 == 0 {
		res = res + "Plang"
	}
	if num%7 == 0 {
		res = res + "Plong"
	}
	if len(res) == 0 {
		res = strconv.Itoa(num)
	}
	return res
}

package railfence

// Action structure
type Action int

const (
	encode Action = iota
	decode
)

// Encode encodes a `msg` string to rail-fence-cipher text.
func Encode(msg string, rails int) string {
	return transform(msg, rails, encode)
}

// Decode decodes a `msg` string to its plain text.
func Decode(msg string, rails int) string {
	return transform(msg, rails, decode)
}

func transform(msg string, nRails int, a Action) string {
	res := []byte(msg)
	var rpos, wpos *int
	var i, j, cycle, delta int
	if a == encode {
		rpos, wpos = &i, &j
	} else {
		rpos, wpos = &j, &i
	}

	cycle = (nRails - 1) * 2
	for rail := 0; rail < nRails; rail++ {
		delta = 2 * rail
		for i = rail; i < len(msg); i += delta {
			res[*wpos] = msg[*rpos]
			j++
			if cycle == delta {
				continue
			}
			delta = cycle - delta
		}
	}
	return string(res)
}

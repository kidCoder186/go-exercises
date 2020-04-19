package robotname

import (
	"errors"
	"math/rand"
	"time"
)

var (
	letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits  = []rune("0123456789")
	nameMap = map[string]bool{}
	allName = 26 * 26 * 10 * 10 * 10
)

// Robot structure
type Robot struct {
	name string
}

func genName() string {
	res := make([]rune, 5)
	res[0] = letters[rand.Intn(len(letters))]
	res[1] = letters[rand.Intn(len(letters))]
	for i := 2; i < 5; i++ {
		res[i] = digits[rand.Intn(len(digits))]
	}
	return string(res)
}

// Name returns current name of the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = genName()
		if len(nameMap) >= allName {
			return "", errors.New("all name are assigned")
		}
		for nameMap[r.name] {
			r.name = genName()
		}
		nameMap[r.name] = true
	}
	return r.name, nil
}

// Reset restarts the robot
func (r *Robot) Reset() {
	rand.Seed(time.Now().UnixNano())
	r.name = ""
}

package perfect

import (
	"errors"
	"math"
)

// Classification structure
type Classification string

const (
	ClassificationAbundant  Classification = "ClassificationAbundant"
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationPerfect   Classification = "ClassificationPerfect"
)

var (
	ErrOnlyPositive = errors.New("negative number")
)

// Classify checks a given number whether is perfect, abundant, or deficient number.
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	var sum int64 = 1
	limit := int64(math.Sqrt(float64(n))) + 1
	for i := int64(2); i < limit; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	if sum == n && sum != 1 {
		return ClassificationPerfect, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}

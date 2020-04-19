package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func count(s string, wg *sync.WaitGroup, c chan FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	c <- m
	wg.Done()
}

// ConcurrentFrequency counts the frequency of each rune in a given text that uses parallelism.
func ConcurrentFrequency(texts []string) FreqMap {
	res := FreqMap{}
	c := make(chan FreqMap, len(texts))
	wg := sync.WaitGroup{}
	for _, text := range texts {
		wg.Add(1)
		go count(text, &wg, c)
	}
	wg.Wait()
	close(c)
	for m := range c {
		for k, v := range m {
			res[k] += v
		}
	}
	return res
}

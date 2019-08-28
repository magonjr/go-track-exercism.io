package letter

// ConcurrentFrequency counts  using goroutines the frequency of each rune in a given text and returns this
// data as a FreqMap.
func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap)

	for _, s := range ss {
		go func(t string) {
			c <- Frequency(t)
		}(s)
	}

	//wait for all goroutines to finish
	for range ss {
		mx := <-c
		for k, v := range mx {
			m[k] += v
		}
	}

	return m
}

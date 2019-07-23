package accumulate

//Accumulate func calcs the accumulate exercice
func Accumulate(values []string, converter func(string) string) []string {
	result := make([]string, len(values))

	for i, v := range values {
		result[i] = converter(v)
	}

	return result
}

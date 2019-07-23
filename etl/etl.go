package etl

import s "strings"

//Transform Extract-Transform-Load (ETL)
func Transform(input map[int][]string) map[string]int {
	if input == nil {
		return nil
	}

	result := make(map[string]int, 23)

	for k, v := range input {
		for l := 0; l < len(v); l++ {
			result[s.ToLower(v[l])] = k
		}
	}

	return result
}

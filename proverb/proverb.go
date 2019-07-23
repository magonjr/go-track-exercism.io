package proverb

import "fmt"

// Proverb exercice solution
func Proverb(rhyme []string) []string {
	n := len(rhyme)

	if n == 0 {
		return []string{}
	}

	proverbs := make([]string, n)

	for i := 0; i < n-1; i++ {
		proverbs[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}

	proverbs[n-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return proverbs
}

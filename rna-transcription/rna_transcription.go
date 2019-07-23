package strand

//ToRNA transforms a DNA sequence in the correspondent RNA sequence
func ToRNA(dna string) string {
	rna := []rune(dna)

	toRnaMap := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

	for i := 0; i < len(rna); i++ {
		rna[i] = toRnaMap[rune(rna[i])]
	}

	return string(rna)
}

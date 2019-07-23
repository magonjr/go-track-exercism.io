package protein

type errStopT struct{}
type errInvalidBaseT struct{}

func (errStopT) Error() string {
	return "ErrStop"
}

func (errInvalidBaseT) Error() string {
	return "ErrInvalidBase"
}

var ErrStop = errStopT{}
var ErrInvalidBase = errInvalidBaseT{}

var aminoMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

//FromRNA converts a RNA sequence in several aminoacids
func FromRNA(rna string) ([]string, error) {
	result := make([]string, 0, 10)
	for i := 0; i < len(rna); i += 3 {
		amin, err := FromCodon(rna[i : i+3])
		if err != nil {
			switch err.(type) {
			case errStopT:
				return result, nil
			case errInvalidBaseT:
				return result, err

			}
		}
		result = append(result, amin)
	}
	return result, nil
}

//FromCodon trasnslates a RNA seq to an aminoacid
func FromCodon(seq string) (string, error) {
	val, exists := aminoMap[seq]
	if !exists {
		return "", ErrInvalidBase
	}

	if val == "STOP" {
		return "", ErrStop
	}
	return val, nil
}

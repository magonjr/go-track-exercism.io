package listops

//IntList typoe for a lit of integers
type IntList []int

//Length returns the length of the list
func (l IntList) Length() int {
	return len(l)
}

//Reverse reverses the list
func (l IntList) Reverse() IntList {
	if l == nil {
		return nil
	}

	aux := l
	for i := 0; i < l.Length(); i++ {
		aux[i] = l[l.Length()-i-1]
	}

	return aux
}

//Append add elements to a list
func (l IntList) Append(other IntList) IntList {
	aux := make(IntList, len(l)+len(other))
	for _, x := range other {
		aux = append(aux, x)
	}
	return aux
}

//Concat concats several lists
func (l IntList) Concat(others []IntList) IntList {
	aux := make(IntList, 0)
	for _, other := range others {
		aux = l.Append(other)
	}
	return aux
}

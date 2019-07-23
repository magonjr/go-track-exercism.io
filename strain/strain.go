package strain

//Ints type
type Ints []int

//Lists type
type Lists [][]int

//Strings type
type Strings []string

//Keep func
func (v Ints) Keep(p func(int) bool) Ints {
	if v == nil {
		return nil
	}
	new := make(Ints, 0, len(v))
	for _, value := range v {
		if p(value) {
			new = append(new, value)
		}
	}
	return new
}

//Discard func
func (v Ints) Discard(p func(int) bool) Ints {
	np := func(v int) bool {
		return !p(v)
	}
	return v.Keep(np)
}

//Keep int slice func
func (v Lists) Keep(p func([]int) bool) Lists {
	if v == nil {
		return nil
	}
	new := make(Lists, 0, len(v))
	for _, value := range v {
		if p(value) {
			new = append(new, value)
		}
	}
	return new
}

//Keep string func
func (v Strings) Keep(p func(string) bool) Strings {
	if v == nil {
		return nil
	}
	new := make(Strings, 0, len(v))
	for _, value := range v {
		if p(value) {
			new = append(new, value)
		}
	}
	return new
}

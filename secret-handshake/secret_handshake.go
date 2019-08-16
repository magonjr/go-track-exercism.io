package secret

import "fmt"

//Handshake
func Handshake(v uint) []string {

	hs := make([]string, 0, 4)

	if v&1 > 0 {
		hs = append(hs, "wink")
	}
	if v&2 > 0 {
		hs = append(hs, "double blink")
	}
	if v&4 > 0 {
		hs = append(hs, "close your eyes")
	}
	if v&8 > 0 {
		hs = append(hs, "jump")
	}
	if v&16 > 0 {
		for i := 0; i < len(hs)/2; i++ {
			hs[i], hs[len(hs)-1-i] = hs[len(hs)-1-i], hs[i]
		}
	}

	fmt.Println(v, hs)

	return hs
}

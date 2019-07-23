// Package twofer exercice
package twofer

// ShareWith func
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

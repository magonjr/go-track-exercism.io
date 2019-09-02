package brackets

var match = map[rune]rune{'[': ']', '(': ')', '{': '}'}
var closing = map[rune]rune{']': '[', ')': '(', '}': '{'}

// Bracket function: Given a string containing brackets `[]`, braces `{}`, parentheses `()`,
// or any combination thereof, verify that any and all pairs are matched
// and nested correctly.
func Bracket(s string) bool {
	queue := make([]rune, 0, 100)
	for _, r := range s {
		if _, ok := match[r]; ok {
			queue = append(queue, r)
			continue
		}

		if _, ok := closing[r]; ok {
			if len(queue) == 0 {
				return false
			}
			top := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if closing[r] != top {
				return false
			}
		}
	}

	return len(queue) == 0
}

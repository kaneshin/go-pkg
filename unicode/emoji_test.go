package unicode

import (
	"testing"
)

// TestIsEmoji ...
func TestIsEmoji(t *testing.T) {

	check := func(s string, expect bool) {
		if expect != func() bool {
			for _, r := range []rune(s) {
				if IsEmoji(r) {
					return true
				}
			}
			return false
		}() {
			t.Fatal("Not Expected")
		}
	}

	check("Hello!", false)
	check("Hello! 😁", true)

}

package parentheses

import (
	"strings"
)

func valid(str string) bool {
	newStr := strings.ReplaceAll(str, "()", "")
	for str != newStr && len(newStr) != 0 {
		str = newStr
		newStr = strings.ReplaceAll(str, "()", "")
	}

	return len(newStr) == 0
}

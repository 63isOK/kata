package parentheses

import (
	"regexp"
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

func validFor(str string) bool {
	level := 0

	for _, x := range str {
		if x == '(' {
			level++
		} else {
			level--
		}
		if level < 0 {
			return false
		}
	}

	return level == 0
}

func validRegexp(str string) bool {
	_, err := regexp.Compile(str)

	return err == nil
}

func validSwitch(s string) bool {
	d := 0
	for _, c := range s {
		switch c {
		case '(':
			d++
			break
		case ')':
			d--
			break
		}
		if d < 0 {
			return false
		}
	}
	return d == 0
}

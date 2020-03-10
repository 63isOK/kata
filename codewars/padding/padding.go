package padding

import (
	"fmt"
	"strconv"
)

func padding(i, w int) string {
	flag := fmt.Sprintf("%%0%ds", w)
	return fmt.Sprintf(flag, strconv.Itoa(i))
}

func paddingNoStd(i, w int) string {
	buf := [20]byte{}
	index := len(buf) - 1

	for i > 0 || w > 0 {
		w--
		c := byte('0' + i%10)
		buf[index] = c

		index--
		i /= 10
	}

	return string(buf[index+1:])
}

func paddingNoStd2(i, w int) string {
	buf := [20]byte{}
	index := len(buf) - 1

	for i > 0 || w > 0 {
		w--
		buf[index] = byte('0' + i%10)

		index--
		i /= 10
	}

	return string(buf[index+1:])
}

func paddingStd(i, w int) string {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || w > 1 {
		w--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)

	return string(b[bp:])
}

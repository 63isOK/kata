package padding

import (
	"fmt"
	"strconv"
)

func padding(i, w int) []byte {
	flag := fmt.Sprintf("%%0%ds", w)
	return []byte(fmt.Sprintf(flag, strconv.Itoa(i)))
}

package shortener

import (
	"fmt"
)

// Make shortened link, also returns the full path because it's easier to use like that.
func Make(link string) string {
	return hash(link)
}

// hash the string using djb2
func hash(s string) string {
	var h uint64 = 5381
	bytes := []byte(s)

	for _, c := range bytes {
		h = ((h << 5) + h) + uint64(c)
	}
	hex := fmt.Sprintf("%x", h)

	return fmt.Sprintf("%s", hex)
}

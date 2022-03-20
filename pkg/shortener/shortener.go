package shortener

import (
	"fmt"
)

const routerPath = "localhost:3000/u/"

// Make shortened link, also returns the full path because it's easier to use like that.
func Make(link string) (string, string) {
	hashed := hash(link)
	result := routerPath + hashed

	return hashed, result
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

package shortener

import (
	"fmt"
)

const routerPath = "localhost:3000/unshort/"

func Make(link string) (string, string) {
	hashed := Hash(link)
	result := routerPath + hashed

	return hashed, result
}

func Hash(s string) string {
	var h uint64 = 5381
	bytes := []byte(s)

	for _, c := range bytes {
		h = ((h << 5) + h) + uint64(c)
	}
	hex := fmt.Sprintf("%x", h)

	return fmt.Sprintf("%s", hex)
}

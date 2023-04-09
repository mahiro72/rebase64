package main

import (
	"fmt"
	"os"

	b64 "github.com/mahiro72/rebase64"
)

func main() {
	if len(os.Args) != 2 {
		panic("invalid args error, args length is not 2")
	}

	inp := os.Args[1]
	fmt.Println("Encoded:", b64.StdEncoding.EncodeToString([]byte(inp)))
}

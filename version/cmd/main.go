package main

import (
	"fmt"
	"os"

	version ".."
)

func main() {
	fmt.Printf("%s", version.Version)

	os.Exit(0)
}

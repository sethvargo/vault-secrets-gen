package main

import (
	"fmt"
	"os"

	version "github.com/sethvargo/vault-secrets-gen/version"
)

func main() {
	fmt.Printf("%s", version.Version)

	os.Exit(0)
}

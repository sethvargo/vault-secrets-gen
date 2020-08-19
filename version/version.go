// Package version defines the version information.
package version

import "fmt"

const Name = "vault-secrets-gen"

var (
	Version   string
	GitCommit string

	HumanVersion = fmt.Sprintf("%s v%s (%s)", Name, Version, GitCommit)
)

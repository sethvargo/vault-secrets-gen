package version

import "fmt"

// Version is the version of the project.
const Version = "0.0.1"

var (
	// Name is the name of the project.
	Name string

	// GitCommit is the git sha of this build.
	GitCommit string

	// HumanVersion is the human-friendly version.
	HumanVersion = fmt.Sprintf("%s v%s (%s)", Name, Version, GitCommit)
)

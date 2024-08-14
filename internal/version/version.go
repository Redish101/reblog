package version

import "fmt"

var (
	Version string = "dev"
	Commit  string = "dev"
)

func GetAppName() string {
	return fmt.Sprintf("acmeidc-%s.%s", Version, Commit)
}

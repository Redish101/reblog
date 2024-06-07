package config

import "fmt"

var (
	Version string = "dev"
	Commit  string = "dev"
)

func GetAppName() string {
	return fmt.Sprintf("reblog-%s.%s", Version, Commit)
}

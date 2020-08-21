package version

import (
	"fmt"
	"runtime"
)

var (
	version = "dev"
	date    = ""
)

func Version() string {
	return fmt.Sprintf("starport version %s %s/%s -build date: %s",
		version,
		runtime.GOOS,
		runtime.GOARCH,
		date)
}

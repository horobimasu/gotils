package gotils

import (
	"os"
	"regexp"
	"strings"
)

var envPathPattern *regexp.Regexp = regexp.MustCompile("%([^%]+)%")

// os.Expand() only expands $... syntax paths
// this expands the BETTER %...% paths
// also ~ -> %USERPROFILE% (only if its the first char ofc)
func ExpandPath(path string) string {
	if []rune(path)[0] == '~' {
		path = strings.Replace(path, "~", "%USERPROFILE%", 1)
	}

	result := envPathPattern.ReplaceAllStringFunc(path, func(match string) string {
		key := match[1 : len(match)-1]

		val := os.Getenv(key)
		if val != "" {
			return val
		}

		return match
	})

	return result
}

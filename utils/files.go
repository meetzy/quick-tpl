package utils

import (
	"os"
	"strings"
)

func JoinSeparator(str ...string) string {
	return strings.Join(str, string(os.PathSeparator))
}
func CheckPath(path *string) {
	if !strings.HasSuffix(*path, string(os.PathSeparator)) {
		*path = *path + string(os.PathSeparator)
	}
}

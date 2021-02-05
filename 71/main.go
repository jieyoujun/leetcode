package simplifypath

import (
	"strings"
)

func simplifyPath(path string) string {
	var paths []string
	for _, p := range strings.Split(path, "/") {
		switch p {
		case "":
			break
		case ".":
			break
		case "..":
			if np := len(paths); np > 0 {
				paths = paths[:np-1]
			}
		default:
			paths = append(paths, p)
		}
	}
	return "/" + strings.Join(paths, "/")
}

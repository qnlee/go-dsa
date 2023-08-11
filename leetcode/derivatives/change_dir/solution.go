package change_dir

import (
	"fmt"
	"strings"
)

/*
Write a cd() function that provides change directory function for an
abstract file system.

Similar to Unix FS:
- Root path is "/"
- Path separator is "/"
- Repeated separators should collapse to a single ("x///y" becomes "x/y")
- Self directory is addressable as "."
- Parent directory is addressable as ".."
- Directory names consist only of English alphabet letters (A-Z and a-z)
- The function should support both relative and absolute paths
*/

func cd(fromPath, toPath string) string {
	if len(toPath) == 1 {
		switch toPath {
		case ".":
			return fromPath
		case "/":
			return toPath
		default:
			return fmt.Sprintf("%s/%s", fromPath, toPath)
		}
	}

	paths := strings.Split(toPath, "/")
	res := strings.Split(fromPath, "/")
	if paths[0] == "" {
		res = []string{""}
	}

	for _, dir := range paths {
		switch dir {
		case "", ".":
			continue
		case "..":
			if len(res) > 1 {
				res = res[:len(res)-1]
			}
		default:
			res = append(res, dir)
		}
	}

	if len(res) == 1 {
		return "/"
	}
	return strings.Join(res, "/")
}

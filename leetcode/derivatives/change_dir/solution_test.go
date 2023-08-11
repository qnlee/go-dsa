package change_dir

import (
	"fmt"
	"testing"
)

func TestCd(t *testing.T) {
	testCases := []struct {
		startingPath   string
		newPath        string
		expectedResult string
	}{
		{"/a/b/c", ".", "/a/b/c"},
		{"/a/b/c", "d/e", "/a/b/c/d/e"},
		{"/a/b/c", "d/./e", "/a/b/c/d/e"},
		{"/a/b/c", "d///./e", "/a/b/c/d/e"},
		{"/a/b/c", "..", "/a/b"},
		{"/a/b/c", "../x", "/a/b/x"},
		{"/a/b/c", "../../1/2", "/a/1/2"},
		{"/a/b/c", "../3/../4/.././5", "/a/b/5"},
		{"/a/b/c", "../../../../..", "/"},
		{"/a/b/c", "/", "/"},
		{"/a/b/c", "/x/y", "/x/y"},
		{"/a/b/c", "/x/y/./../z///.", "/x/z"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("start:%s  cd to:%s", tc.startingPath, tc.newPath), func(t *testing.T) {
			result := cd(tc.startingPath, tc.newPath)
			if result != tc.expectedResult {
				t.Errorf("got: %s, expected: %s", result, tc.expectedResult)
			}
		})
	}
}

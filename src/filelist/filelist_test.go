/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package filelist

import "testing"

func TestFileList(t *testing.T) {
	if err := DeleteFileAndDir("/home/lbbxsxlz/TestDir"); err != nil {
		t.Error("delete files fail")
	}
}

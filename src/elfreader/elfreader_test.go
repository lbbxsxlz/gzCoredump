/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package elfreader

import "testing"

func TestElfReader(t *testing.T) {
	if err := CreateUuidFile("/home/lbbxsxlz/BigBang", ".uuid", "BigBang-uuid.txt"); err != nil {
		t.Error("CreateUuidFile fail")
	}
}


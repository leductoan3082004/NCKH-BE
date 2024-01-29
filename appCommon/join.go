package appCommon

import "strings"

func Join(sep string, val ...string) string {
	return strings.Join(val, sep)
}

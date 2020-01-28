package lib

import "strings"

// TransferString for sqli
func TransferString(str string) string {
	str = strings.ReplaceAll(str, "\\", "\\\\'")
	return strings.ReplaceAll(str, "'", "\\'")
}

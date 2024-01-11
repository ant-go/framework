package ierror

import (
	"fmt"
	"strings"
)

func fieldsBuilder(slice []any) string {
	var sb strings.Builder
	for i, item := range slice {
		if i%2 == 0 {
			_, _ = sb.WriteString(fmt.Sprintf("%v=", item))
		} else {
			_, _ = sb.WriteString(fmt.Sprintf("%v|", item))
		}
	}
	return sb.String()
}


package main

import (
    "runtime"
	"strings"
)


func ReceiverName(n int) string {
	pc, _, _, _ := runtime.Caller(n)
	f := runtime.FuncForPC(pc)
	parts := strings.Split(f.Name(), ".")
	if len(parts) < 2 {
		return "Couldn't not get receiver name"
	}
	return parts[1]
}



















package checkers

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

var Keywords = []string{"OK", "WARNING", "CRITICAL", "UNKNOWN"}

func StrToStatus(s string) Status {
	switch s {
	case "OK":
		return OK
	case "WARNING":
		return WARNING
	case "CRITICAL":
		return CRITICAL
	case "UNKNOWN":
		return UNKNOWN
	}
	panic("invalid inputs.")
}
// Parse parses <status>=<status> notation string. <status> is one of:
//   - ok
//   - warning
//   - critical
//   - unknown
//
// It can have multiple key-value pairs by comma.
func Parse(arg string) (map[Status]Status, error) {
	if arg == "" {
		return nil, nil
	}

	maps := make(map[Status]Status, 0)
	args := strings.Split(arg, ",")

	for _, s := range args {
		if s == "" {
			continue
		}
		v := strings.Split(s, "=")
		if len(v) != 2 {
			return nil, fmt.Errorf("arguments is invalid : %v", v)
		}
		from := strings.ToUpper(v[0])
		to := strings.ToUpper(v[1])
		if !slices.Contains(Keywords, from) {
			return nil, fmt.Errorf("from argument is invalid : %s", v[0])
		}
		if !slices.Contains(Keywords, to) {
			return nil, fmt.Errorf("to argument is invalid : %s", v[1])
		}
		if _, ok := maps[StrToStatus(from)]; ok {
			return nil, fmt.Errorf("arguments is duplicate : %v", from)
		}
		maps[StrToStatus(from)] = StrToStatus(to)
	}
	return maps, nil
}

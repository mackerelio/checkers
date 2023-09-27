package checkers

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

var keywords = []string{"OK", "WARNING", "CRITICAL", "UNKNOWN"}

func strToStatus(s string) Status {
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
		rawFrom, rawTo, found := strings.Cut(s, "=")
		if !found {
			return nil, fmt.Errorf("arguments is invalid : %v", s)
		}
		from := strings.ToUpper(rawFrom)
		to := strings.ToUpper(rawTo)
		if !slices.Contains(keywords, from) {
			return nil, fmt.Errorf("from argument is invalid : %s", rawFrom)
		}
		if !slices.Contains(keywords, to) {
			return nil, fmt.Errorf("to argument is invalid : %s", rawTo)
		}
		if _, ok := maps[strToStatus(from)]; ok {
			return nil, fmt.Errorf("arguments is duplicate : %v", from)
		}
		maps[strToStatus(from)] = strToStatus(to)
	}
	return maps, nil
}

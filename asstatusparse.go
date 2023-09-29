package checkers

import "strings"

type invalidStatusMapError string

func (err invalidStatusMapError) Error() string {
	return "invalid status map: " + string(err)
}

type duplicateStatusMapError Status

func (err duplicateStatusMapError) Error() string {
	return "duplicate status in map: " + Status(err).String()
}

func parseStatusPrefix(src string) (Status, string, bool) {
	switch {
	case strings.HasPrefix(src, "OK"):
		return OK, src[len("OK"):], true
	case strings.HasPrefix(src, "WARNING"):
		return WARNING, src[len("WARNING"):], true
	case strings.HasPrefix(src, "CRITICAL"):
		return CRITICAL, src[len("CRITICAL"):], true
	case strings.HasPrefix(src, "UNKNOWN"):
		return UNKNOWN, src[len("UNKNOWN"):], true
	default:
		return OK, src, false
	}
}

// Parse parses a string of the form <status>=<status>. <status> is one of:
//
//   - ok
//   - warning
//   - critical
//   - unknown
//
// You can specify multiple key-value pairs separated by commas.
func Parse(src string) (map[Status]Status, error) {
	orig, stm := src, map[Status]Status{}
	for src = strings.ToUpper(src); src != ""; {
		var from, to Status
		var ok bool
		if from, src, ok = parseStatusPrefix(src); !ok {
			return nil, invalidStatusMapError(orig)
		}
		if src, ok = strings.CutPrefix(src, "="); !ok {
			return nil, invalidStatusMapError(orig)
		}
		if to, src, ok = parseStatusPrefix(src); !ok {
			return nil, invalidStatusMapError(orig)
		}
		if src, ok = strings.CutPrefix(src, ","); ok == (src == "") {
			return nil, invalidStatusMapError(orig)
		}
		if _, ok = stm[from]; ok {
			return nil, duplicateStatusMapError(from)
		}
		stm[from] = to
	}
	return stm, nil
}

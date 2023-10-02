package checkers

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		expected map[Status]Status
		wantErr  bool
	}{
		{
			name:   "normal",
			source: "unknown=critical",
			expected: map[Status]Status{
				UNKNOWN: CRITICAL,
			},
		},
		{
			name:   "uppercase",
			source: "UNKNOWN=CRITICAL",
			expected: map[Status]Status{
				UNKNOWN: CRITICAL,
			},
		},
		{
			name:   "multi",
			source: "unknown=ok,warning=critical",
			expected: map[Status]Status{
				UNKNOWN: OK,
				WARNING: CRITICAL,
			},
		},
		{
			name:    "from error",
			source:  "invalid=critical",
			wantErr: true,
		},
		{
			name:    "= error",
			source:  "unknown critical",
			wantErr: true,
		},
		{
			name:    "to error",
			source:  "critical=invalid",
			wantErr: true,
		},
		{
			name:    "argument error",
			source:  "unknown=critical=invalid",
			wantErr: true,
		},
		{
			name:    "missing comma error",
			source:  "unknown=criticalwarning=critical",
			wantErr: true,
		},
		{
			name:    "trailing comma error",
			source:  "unknown=ok,",
			wantErr: true,
		},
		{
			name:    "duplicate error",
			source:  "critical=unknown,critical=warning",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Parse(tt.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("Parse() should be %+v but got: %+v", tt.expected, actual)
			}
		})
	}
}

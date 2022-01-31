package ismgExporter

import (
	"fmt"
	"testing"
)

func TestPort_Create(t *testing.T) {
	tests := []struct {
		port    int
		success bool
	}{
		{0, false},
		{999, false},
		{1000, false},
		{1001, true},
		{8080, true},
		{65535, true},
		{65536, false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.port), func(t *testing.T) {
			valid := isPortValid(&tc.port)
			if valid != tc.success {
				t.Errorf("Verification with port value: %d, was failed: expected result of validate is: %v but result is: %v", tc.port, tc.success, valid)
			}
		})
	}
}

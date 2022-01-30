package ismgExporter

import (
	"testing"
)

func TestPort_Create(t *testing.T) {
	PortTesting(0, false, t)
	PortTesting(999, false, t)
	PortTesting(1000, false, t)
	PortTesting(1001, true, t)
	PortTesting(8080, true, t)
	PortTesting(65535, true, t)
	PortTesting(65536, false, t)
}

func PortTesting(value int, success bool, t *testing.T) {
	valid := isPortValid(&value)
	if valid != success {
		t.Errorf("Verification with port value: %d, was failed: expected result of validate is: %v but result is: %v", value, success, valid)
	}

}

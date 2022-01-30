package ismgExporter_test

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
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
	portStruct := new(ismgExporter.Port)
	_, err := portStruct.Create(value)
	if (err != nil) == success {
		t.Errorf("Verification with port value: %d, was failed: expected result of success is: %v", value, success)
	}

}

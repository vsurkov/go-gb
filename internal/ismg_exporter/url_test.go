package ismgExporter_test

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"testing"
)

func TestURL_Create(t *testing.T) {
	URLTesting("postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", true, t)
	URLTesting("postgres://db-user:db-password@petstore-db:5432", true, t)
	URLTesting("postgres://db-user@petstore-db:5432/petstore?sslmode=disable", true, t)
	URLTesting("postgres://petstore-db:5432/petstore?sslmode=disable", true, t)
	URLTesting("http://jaeger:16686", true, t)
	URLTesting("http://sentry:9000", true, t)
	URLTesting("kafka:9000", true, t)
	//todo fix bugs with IP address verification
	//URLTesting("127.0.0.1", true, t)
	//URLTesting("127.0.0.1:8080", true, t)
	URLTesting("foo", false, t)
	URLTesting("http://", false, t)
	URLTesting("9090", false, t)
	//todo fix bug with port
	//URLTesting("http://9090", false, t)
}

func URLTesting(value string, success bool, t *testing.T) {
	urlStruct := new(ismgExporter.URL)
	_, err := urlStruct.Create(value)
	if (err != nil) == success {
		t.Errorf("Verification with URL value: %s, was failed: expected result of success is: %v", value, success)
	}

}

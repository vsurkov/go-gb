package ismgExporter

import (
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
	URLTesting("127.0.0.1:8080", true, t)

	URLTesting("postgres://db-user:db-password@petstore-db:100000/petstore?sslmode=disable", false, t) // wrong port
	URLTesting("http://jaeger:100", false, t)                                                          // wrong port
	URLTesting("http://jaeger:100000", false, t)                                                       // wrong port
	URLTesting("127.0.0.1", false, t)                                                                  // no port
	URLTesting("foo", false, t)                                                                        // no port
	URLTesting("http://", false, t)                                                                    // no host
	URLTesting("9090", false, t)                                                                       // no port
	URLTesting("http://9090", false, t)                                                                // no port
}

func URLTesting(value string, success bool, t *testing.T) {
	valid := isUrlValid(&value)
	if valid != success {
		t.Errorf("Verification with URL value: %s, was failed: expected result of validate is: %v but result is: %v", value, success, valid)
	}

}

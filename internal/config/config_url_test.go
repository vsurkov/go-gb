package ismgExporter

import (
	"testing"
)

func TestURL_Create(t *testing.T) {

	tests := []struct {
		url     string
		success bool
	}{
		{"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", true},
		{"postgres://db-user:db-password@petstore-db:5432", true},
		{"postgres://db-user@petstore-db:5432/petstore?sslmode=disable", true},
		{"postgres://petstore-db:5432/petstore?sslmode=disable", true},
		{"http://jaeger:16686", true},
		{"http://sentry:9000", true},
		{"kafka:9000", true},
		{"127.0.0.1:8080", true},

		{"postgres://db-user:db-password@petstore-db:100000/petstore?sslmode=disable", false}, // wrong port
		{"http://jaeger:100", false},    // wrong port
		{"http://jaeger:100000", false}, // wrong port
		{"127.0.0.1", false},            // no port
		{"foo", false},                  // no port
		{"http://", false},              // no host
		{"9090", false},                 // no port
		{"http://9090", false},          // no port
	}

	for _, tc := range tests {
		t.Run(tc.url, func(t *testing.T) {
			valid := isUrlValid(&tc.url)
			if valid != tc.success {
				t.Errorf("Verification with URL value: %s, was failed: expected result of validate is: %v but result is: %v", tc.url, tc.success, valid)
			}
		})
	}

}

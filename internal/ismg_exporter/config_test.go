package ismgExporter_test

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"testing"
)

func TestNewConfig(t *testing.T) {

	expected := ismgExporter.Config{
		Port:         ismgExporter.Port{9010},
		DB_url:       ismgExporter.URL{"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"},
		Jaeger_url:   ismgExporter.URL{"http://jaeger:16686"},
		Sentry_url:   ismgExporter.URL{"http://sentry:9000"},
		Kafka_broker: ismgExporter.URL{"kafka:9000"},
		App_id:       "DEFAULT_APP_ID",
		App_key:      "8787928792847928749248742479724",
	}
	received := new(ismgExporter.Config)
	received.Load()

	checkIntField(expected.Port.Value, received.Port.Value, t)
	checkStrField(expected.DB_url.Value, received.DB_url.Value, t)
	checkStrField(expected.Jaeger_url.Value, received.Jaeger_url.Value, t)
	checkStrField(expected.Sentry_url.Value, received.Sentry_url.Value, t)
	checkStrField(expected.Kafka_broker.Value, received.Kafka_broker.Value, t)
	checkStrField(expected.App_id, received.App_id, t)
	checkStrField(expected.App_key, received.App_key, t)

}

func checkIntField(expected int, received int, t *testing.T) {
	if expected != received {
		t.Errorf("Expected %d, but received %d", expected, received)
	}
}

func checkStrField(expected string, received string, t *testing.T) {
	if expected != received {
		t.Errorf("Expected %s, but received %s", expected, received)
	}
}

package ismgExporter_test

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"testing"
)

func TestNewConfig(t *testing.T) {

	expected := ismgExporter.Config{
		Port:        ismgExporter.Port{9010},
		DbURL:       ismgExporter.URL{"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"},
		JaegerURL:   ismgExporter.URL{"http://jaeger:16686"},
		SentryURL:   ismgExporter.URL{"http://sentry:9000"},
		KafkaBroker: ismgExporter.URL{"kafka:9000"},
		AppID:       "DEFAULT_APP_ID",
		AppKey:      "8787928792847928749248742479724",
	}
	received := new(ismgExporter.Config)
	received.LoadFromParams()

	checkIntField(expected.Port.Value, received.Port.Value, t)
	checkStrField(expected.DbURL.Value, received.DbURL.Value, t)
	checkStrField(expected.JaegerURL.Value, received.JaegerURL.Value, t)
	checkStrField(expected.SentryURL.Value, received.SentryURL.Value, t)
	checkStrField(expected.KafkaBroker.Value, received.KafkaBroker.Value, t)
	checkStrField(expected.AppID, received.AppID, t)
	checkStrField(expected.AppKey, received.AppKey, t)
}

func TestFileConfig(t *testing.T) {

	expected := ismgExporter.Config{
		Port:        ismgExporter.Port{9011},
		DbURL:       ismgExporter.URL{"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=enable"},
		JaegerURL:   ismgExporter.URL{"http://jaeger:16686"},
		SentryURL:   ismgExporter.URL{"http://sentry:9000"},
		KafkaBroker: ismgExporter.URL{"kafka:9000"},
		AppID:       "DEFAULT_APP_ID",
		AppKey:      "8787928792847928749248742479724",
	}
	received := new(ismgExporter.Config)
	received.LoadFromConfig("..//..//configs//ismg_exporter.json")

	checkIntField(expected.Port.Value, received.Port.Value, t)
	checkStrField(expected.DbURL.Value, received.DbURL.Value, t)
	checkStrField(expected.JaegerURL.Value, received.JaegerURL.Value, t)
	checkStrField(expected.SentryURL.Value, received.SentryURL.Value, t)
	checkStrField(expected.KafkaBroker.Value, received.KafkaBroker.Value, t)
	checkStrField(expected.AppID, received.AppID, t)
	checkStrField(expected.AppKey, received.AppKey, t)

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

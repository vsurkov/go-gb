package ismgExporter_test

import (
	"fmt"
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"testing"
)

func TestNewConfig(t *testing.T) {
	expected := ismgExporter.Config{
		Port:        9015,
		DbUrl:       "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=enable",
		JaegerUrl:   "http://jaeger:16687",
		SentryUrl:   "http://sentry:9001",
		KafkaBroker: "kafka:9001",
		AppId:       "app_id_for_testing",
		AppKey:      "app_key_for_testing",
	}
	received := ismgExporter.Config{}
	received.Load("..//..//configs//ismg_exporter.json")

	checkIntField(expected.Port, received.Port, t)
	checkStrField(expected.DbUrl, received.DbUrl, t)
	checkStrField(expected.JaegerUrl, received.JaegerUrl, t)
	checkStrField(expected.SentryUrl, received.SentryUrl, t)
	checkStrField(expected.KafkaBroker, received.KafkaBroker, t)
	checkStrField(expected.AppId, received.AppId, t)
	checkStrField(expected.AppKey, received.AppKey, t)
	fmt.Printf("TEST received config:%v", received)
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

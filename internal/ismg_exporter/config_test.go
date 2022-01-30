package ismgExporter_test

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"log"
	"testing"
)

func TestNewConfig(t *testing.T) {
	expected := ismgExporter.Config{
		Port:        9010,
		DbUrl:       "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
		JaegerUrl:   "http://jaeger:16686",
		SentryUrl:   "http://sentry:9000",
		KafkaBroker: "kafka:9000",
		AppId:       "app_id_for_testing",
		AppKey:      "app_key_for_testing",
	}
	received := ismgExporter.Config{}
	received.Load()

	checkIntField(expected.Port, received.Port, t)
	checkStrField(expected.DbUrl, received.DbUrl, t)
	checkStrField(expected.JaegerUrl, received.JaegerUrl, t)
	checkStrField(expected.SentryUrl, received.SentryUrl, t)
	checkStrField(expected.KafkaBroker, received.KafkaBroker, t)
	checkStrField(expected.AppId, received.AppId, t)
	checkStrField(expected.AppKey, received.AppKey, t)
}

func TestFileConfig(t *testing.T) {

	expected := ismgExporter.Config{
		Port:        9010,
		DbUrl:       "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable",
		JaegerUrl:   "http://jaeger:16686",
		SentryUrl:   "http://sentry:9000",
		KafkaBroker: "kafka:9000",
		AppId:       "app_id_for_testing",
		AppKey:      "app_key_for_testing",
	}
	received := ismgExporter.Config{}
	err := received.LoadFromFile("..//..//configs//ismg_exporter.json")
	if err != nil {
		log.Printf("Error on loadign from File %v", err)
	}

	checkIntField(expected.Port, received.Port, t)
	checkStrField(expected.DbUrl, received.DbUrl, t)
	checkStrField(expected.JaegerUrl, received.JaegerUrl, t)
	checkStrField(expected.SentryUrl, received.SentryUrl, t)
	checkStrField(expected.KafkaBroker, received.KafkaBroker, t)
	checkStrField(expected.AppId, received.AppId, t)
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

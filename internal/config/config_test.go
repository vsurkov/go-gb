package ismgExporter_test

import (
	"github.com/stretchr/testify/assert"
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/config"
	"reflect"
	"testing"
)

func TestPositiveConfig(t *testing.T) {
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

	// Общее сравнение двух объектов
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("DeepEqual fail")
	}

	// Сравнение по полям, позволяет увидеть где именно ошибка
	assert.Equal(t, expected.Port, received.Port, "Values must be equal")
	assert.Equal(t, expected.DbUrl, received.DbUrl, "RValues must be equal")
	assert.Equal(t, expected.JaegerUrl, received.JaegerUrl, "ReceiValues must be equal")
	assert.Equal(t, expected.SentryUrl, received.SentryUrl, "ReceiValues must be equal")
	assert.Equal(t, expected.KafkaBroker, received.KafkaBroker, "ReceiveValues must be equal")
	assert.Equal(t, expected.AppId, received.AppId, "RValues must be equal")
	assert.Equal(t, expected.AppKey, received.AppKey, "ReValues must be equal")
}

func TestNegativeConfig(t *testing.T) {
	expected := ismgExporter.Config{
		Port:        9000,
		DbUrl:       "postgres://db-user:db-password@petstore-db:5432",
		JaegerUrl:   "http://jaeger",
		SentryUrl:   "http://sentry",
		KafkaBroker: "kafka",
		AppId:       "app_id_for_bad_testing",
		AppKey:      "app_key_for_bad_testing",
	}
	received := ismgExporter.Config{}
	received.Load("..//..//configs//ismg_exporter.json")

	// Общее сравнение двух объектов
	if reflect.DeepEqual(expected, received) {
		t.Errorf("DeepEqual fail objects can't be equal")
	}

	// Сравнение по полям, позволяет увидеть где именно ошибка
	assert.NotEqual(t, expected.Port, received.Port, "Values can't be equal")
	assert.NotEqual(t, expected.DbUrl, received.DbUrl, "RValues can't be equal")
	assert.NotEqual(t, expected.JaegerUrl, received.JaegerUrl, "ReceiValues can't be equal")
	assert.NotEqual(t, expected.SentryUrl, received.SentryUrl, "ReceiValues can't be equal")
	assert.NotEqual(t, expected.KafkaBroker, received.KafkaBroker, "ReceiveValues can't be equal")
	assert.NotEqual(t, expected.AppId, received.AppId, "RValues can't be equal")
	assert.NotEqual(t, expected.AppKey, received.AppKey, "ReValues can't be equal")
}

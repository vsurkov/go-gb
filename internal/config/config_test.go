package ismgExporter_test

import (
	"github.com/stretchr/testify/assert"
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/config"
	"os"
	"reflect"
	"testing"
)

var (
	expected, expectedFail, received ismgExporter.Config
)

func TestMain(m *testing.M) {
	//Конфигурация совпадающая с тестовым конфигом
	expected = ismgExporter.Config{
		Port:        9015,
		DbUrl:       "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=enable",
		JaegerUrl:   "http://jaeger:16687",
		SentryUrl:   "http://sentry:9001",
		KafkaBroker: "kafka:9001",
		AppId:       "app_id_for_testing",
		AppKey:      "app_key_for_testing",
	}

	// Конфигурация НЕ совпадающая с тестовым конфигом, для fail тестов
	expectedFail = ismgExporter.Config{
		Port:        9000,
		DbUrl:       "postgres://db-user:db-password@petstore-db:5432",
		JaegerUrl:   "http://jaeger",
		SentryUrl:   "http://sentry",
		KafkaBroker: "kafka",
		AppId:       "app_id_for_bad_testing",
		AppKey:      "app_key_for_bad_testing",
	}

	// Ожидаемая конфигурация, загрузка из файлов поддерживаемых типов
	received.Load("../../configs/ismg_exporter.yaml")
	exitVal := m.Run()
	//"Do stuff AFTER the tests!"
	os.Exit(exitVal)
}

func TestPositiveJsonConfig(t *testing.T) {
	// Общее сравнение двух объектов
	if !reflect.DeepEqual(expected, received) {
		t.Errorf("DeepEqual fail")
	}

	// Сравнение по полям, позволяет увидеть где именно ошибка
	assert.Equal(t, expected.Port, received.Port, "Values must be equal")
	assert.Equal(t, expected.DbUrl, received.DbUrl, "Values must be equal")
	assert.Equal(t, expected.JaegerUrl, received.JaegerUrl, "Values must be equal")
	assert.Equal(t, expected.SentryUrl, received.SentryUrl, "Values must be equal")
	assert.Equal(t, expected.KafkaBroker, received.KafkaBroker, "Values must be equal")
	assert.Equal(t, expected.AppId, received.AppId, "Values must be equal")
	assert.Equal(t, expected.AppKey, received.AppKey, "Values must be equal")
}

func TestNegativeJsonConfig(t *testing.T) {
	// Общее сравнение двух объектов
	if reflect.DeepEqual(expectedFail, received) {
		t.Errorf("DeepEqual fail objects can't be equal")
	}

	// Сравнение по полям, позволяет увидеть где именно ошибка
	assert.NotEqual(t, expectedFail.Port, received.Port, "Values can't be equal")
	assert.NotEqual(t, expectedFail.DbUrl, received.DbUrl, "Values can't be equal")
	assert.NotEqual(t, expectedFail.JaegerUrl, received.JaegerUrl, "Values can't be equal")
	assert.NotEqual(t, expectedFail.SentryUrl, received.SentryUrl, "Values can't be equal")
	assert.NotEqual(t, expectedFail.KafkaBroker, received.KafkaBroker, "Values can't be equal")
	assert.NotEqual(t, expectedFail.AppId, received.AppId, "Values can't be equal")
	assert.NotEqual(t, expectedFail.AppKey, received.AppKey, "Values can't be equal")
}

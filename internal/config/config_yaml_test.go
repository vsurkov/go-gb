package ismgExporter_test

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPositiveYamlConfig(t *testing.T) {
	// Общее сравнение двух объектов
	if !reflect.DeepEqual(expected, recYamlСfg) {
		t.Errorf("DeepEqual fail")
	}

	// Сравнение по полям, позволяет увидеть где именно ошибка
	assert.Equal(t, expected.Port, recYamlСfg.Port, "Values must be equal")
	assert.Equal(t, expected.DbUrl, recYamlСfg.DbUrl, "RValues must be equal")
	assert.Equal(t, expected.JaegerUrl, recYamlСfg.JaegerUrl, "ReceiValues must be equal")
	assert.Equal(t, expected.SentryUrl, recYamlСfg.SentryUrl, "ReceiValues must be equal")
	assert.Equal(t, expected.KafkaBroker, recYamlСfg.KafkaBroker, "ReceiveValues must be equal")
	assert.Equal(t, expected.AppId, recYamlСfg.AppId, "RValues must be equal")
	assert.Equal(t, expected.AppKey, recYamlСfg.AppKey, "ReValues must be equal")
}

func TestNegativeYamlConfig(t *testing.T) {
	// Общее сравнение двух объектов
	if reflect.DeepEqual(expectedFail, recYamlСfg) {
		t.Errorf("DeepEqual fail objects can't be equal")
	}

	// Сравнение по полям, позволяет увидеть где именно ошибка
	assert.NotEqual(t, expectedFail.Port, recYamlСfg.Port, "Values can't be equal")
	assert.NotEqual(t, expectedFail.DbUrl, recYamlСfg.DbUrl, "RValues can't be equal")
	assert.NotEqual(t, expectedFail.JaegerUrl, recYamlСfg.JaegerUrl, "ReceiValues can't be equal")
	assert.NotEqual(t, expectedFail.SentryUrl, recYamlСfg.SentryUrl, "ReceiValues can't be equal")
	assert.NotEqual(t, expectedFail.KafkaBroker, recYamlСfg.KafkaBroker, "ReceiveValues can't be equal")
	assert.NotEqual(t, expectedFail.AppId, recYamlСfg.AppId, "RValues can't be equal")
	assert.NotEqual(t, expectedFail.AppKey, recYamlСfg.AppKey, "ReValues can't be equal")
}

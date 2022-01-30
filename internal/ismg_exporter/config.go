package ismgExporter

import (
	"encoding/json"
	"fmt"
	"github.com/goware/urlx"
	"log"
	"os"
	"strconv"
)
import "github.com/namsral/flag"

// Config Описание базовой структуры конфигурации
type Config struct {
	Port        int    `json:"port"`
	DbUrl       string `json:"db_url"`
	JaegerUrl   string `json:"jaeger_url"`
	SentryUrl   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	AppId       string `json:"app_id"`
	AppKey      string `json:"app_key"`
}

// NewConfig Функция возвращает инициализированную пустую структуру Config
func NewConfig() *Config {
	return &Config{
		Port:        0,
		DbUrl:       "",
		JaegerUrl:   "",
		SentryUrl:   "",
		KafkaBroker: "",
		AppId:       "",
		AppKey:      "",
	}
}

// Определяем значения по умолчанию, нужны отдельно что бы к ним вернуться при не валидности полученного
const (
	defaultPort        = 9010
	defaultDBUrl       = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	defaultJaegerUrl   = "http://jaeger:16686"
	defaultSentryUrl   = "http://sentry:9000"
	defaultKafkaBroker = "kafka:9000"
	defaultSomeAppID   = "app_id_for_testing"
	defaultSomeAppKey  = "app_key_for_testing"
)

// Определяем флаги
var (
	flagPort        = flag.Int("port", defaultPort, "Server port, must be in the range 1000-65535")
	flagDBUrl       = flag.String("db_url", defaultDBUrl, "Database connection string, example: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	flagJaegerUrl   = flag.String("jaeger_url", defaultJaegerUrl, "Jaeger URL, example: http://jaeger:16686k")
	flagSentryUrl   = flag.String("sentry_url", defaultSentryUrl, "Sentry URL, example: http://sentry:9000")
	flagKafkaBroker = flag.String("kafka_broker", defaultKafkaBroker, "Kafka broker URI, example: kafka:9092")
	flagSomeAppID   = flag.String("app_id", defaultSomeAppID, "Application ID")
	flagSomeAppKey  = flag.String("app_key", defaultSomeAppKey, "Application KEY")
)

// LoadFromConfig Загружаем конфигурацию из файла
func (config *Config) LoadFromConfig(configPath string) error {
	// Отдельная структура нужна потому что Config содержит в полях струры а не простые типы,
	// для Decode нужно описание своей структуры для загрузки и последующего приведения к Config

	// Открываем файл по configPath и проверяем на ошибки
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("load config failed. %v", err)
	}

	// Закрываем октытый файл в любом случае
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("can't close file. %v", err)
		}
	}()

	// Декодируем конфигурацию и заполняем структуру конфигурации
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		fmt.Errorf("can't decode JSON configuration on configPath: %v. Error: \n %v", configPath, err)
	}

	return nil
}

// LoadFromParams Парсим флаги и загружаем через функции загрузки параметры конфигурации
func (config *Config) LoadFromParams() {
	flag.Parse()
	// Каждый из параметров заполняем результатом загрузки параметра
	config.Port = *loadPortParam(flagPort, defaultPort)
	config.DbUrl = *loadURLParam(flagDBUrl, defaultDBUrl)
	config.JaegerUrl = *loadURLParam(flagJaegerUrl, defaultJaegerUrl)
	config.SentryUrl = *loadURLParam(flagSentryUrl, defaultSentryUrl)
	config.KafkaBroker = *loadURLParam(flagKafkaBroker, defaultKafkaBroker)
	config.AppId = *loadStringParam(flagSomeAppID, defaultSomeAppID)
	config.AppKey = *loadStringParam(flagSomeAppKey, defaultSomeAppID)
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает Int
func loadPortParam(flag *int, defaultValue int) *int {
	if flag != nil {
		if isPortValid(flag) {
			return flag
		}
		log.Printf("Port %v invalid (must be int and in the range 1000-65535), loaded default value: %d \n", *flag, defaultValue)
		return &defaultValue
	}
	return &defaultValue
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает String
func loadURLParam(flag *string, defaultValue string) *string {
	if flag != nil {
		if isUrlValid(flag) {
			return flag
		}
		log.Printf("url invalid, loaded default value: %v \n", defaultValue)
		return &defaultValue
	}
	return &defaultValue
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает String
func loadStringParam(flag *string, defaultValue string) *string {
	if *flag != "" {
		return flag
	}
	log.Printf("param invalid, loaded default value: %v \n", defaultValue)
	return &defaultValue
}

// Валидирует порт, проверяет, что указанное значение порта находится в разрешенном диапазоне от 1000 до 65535
func isPortValid(value *int) bool {
	if *value <= 1000 || *value > 65535 {
		return false
	}
	return true
}

func isUrlValid(raw *string) bool {
	url, err := urlx.Parse(*raw)
	if err != nil || url.Scheme == "" || url.Host == "" {
		log.Printf("URL %v: %v", *raw, err)
		return false
	}

	// Проверяем валидность порта в url
	_, portStr, _ := urlx.SplitHostPort(url)

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Port %v invalid (must be int and in the range 1000-65535), error: %v \n", port, err)
		return false
	}

	if err != nil || !isPortValid(&port) {
		log.Printf("Port %v invalid (must be int and in the range 1000-65535) \n", port)
		return false
	}
	return true
}

package ismgExporter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
import "github.com/namsral/flag"

// Config Описание базовой структуры конфигурации
type Config struct {
	Port        Port
	DbURL       URL
	JaegerURL   URL
	SentryURL   URL
	KafkaBroker URL
	AppID       string
	AppKey      string
}

// NewConfig Функция возвращает инициализированную пустую структуру Config
func NewConfig() *Config {
	return &Config{
		Port:        Port{0},
		DbURL:       URL{""},
		JaegerURL:   URL{""},
		SentryURL:   URL{""},
		KafkaBroker: URL{""},
		AppID:       "",
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
	defaultSomeAppID   = "DEFAULT_APP_ID"
	defaultSomeAppKey  = "8787928792847928749248742479724"
)

// Определяем флаги
var (
	flagPort = flag.Int(
		"port",
		defaultPort,
		"Server port")
	flagDBUrl = flag.String(
		"db_url",
		defaultDBUrl,
		"Database connection string, example: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	flagJaegerUrl = flag.String(
		"jaeger_url",
		defaultJaegerUrl,
		"Jaeger URL, example: http://jaeger:16686k")
	flagSentryUrl = flag.String(
		"sentry_url",
		defaultSentryUrl,
		"Sentry URL, example: http://sentry:9000")
	flagKafkaBroker = flag.String(
		"kafka_broker",
		defaultKafkaBroker,
		"Kafka broker URI, example: kafka:9092")
	flagSomeAppID = flag.String(
		"app_id",
		defaultSomeAppID,
		"Application ID")
	flagSomeAppKey = flag.String(
		"app_key",
		defaultSomeAppKey,
		"Application KEY")
)

// LoadFromConfig Загружаем конфигурацию из файла
func (config *Config) LoadFromConfig(configPath string) error {
	// Отдельная структура нужна потому что Config содержит в полях струры а не простые типы,
	// для Decode нужно описание своей структуры для загрузки и последующего приведения к Config
	type fileConfig struct {
		Port struct {
			Value int `json:"Value"`
		} `json:"port"`
		DbURL struct {
			Value string `json:"Value"`
		} `json:"db_url"`
		JaegerURL struct {
			Value string `json:"Value"`
		} `json:"jaeger_url"`
		SentryURL struct {
			Value string `json:"Value"`
		} `json:"sentry_url"`
		KafkaBroker struct {
			Value string `json:"Value"`
		} `json:"kafka_broker"`
		AppID  string `json:"app_id"`
		AppKey string `json:"app_key"`
	}

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
	loadedConfig := new(fileConfig)
	err = json.NewDecoder(file).Decode(loadedConfig)
	if err != nil {
		fmt.Errorf("can't decode JSON configuration on configPath: %v. Error: \n %v", configPath, err)
	}

	// Каждый из параметров заполняем результатом загрузки параметра
	config.Port = *loadPortParam(flagPort, loadedConfig.Port.Value)
	config.DbURL = *loadURLParam(flagDBUrl, loadedConfig.DbURL.Value)
	config.JaegerURL = *loadURLParam(flagJaegerUrl, loadedConfig.JaegerURL.Value)
	config.SentryURL = *loadURLParam(flagSentryUrl, loadedConfig.SentryURL.Value)
	config.KafkaBroker = *loadURLParam(flagKafkaBroker, loadedConfig.KafkaBroker.Value)
	config.AppID = *loadStringParam(flagSomeAppID, loadedConfig.AppID)
	config.AppKey = *loadStringParam(flagSomeAppKey, loadedConfig.AppKey)
	return nil
}

// LoadFromParams Парсим флаги и загружаем через функции загрузки параметры конфигурации
func (config *Config) LoadFromParams() {
	flag.Parse()
	// Каждый из параметров заполняем результатом загрузки параметра
	config.Port = *loadPortParam(flagPort, defaultPort)
	config.DbURL = *loadURLParam(flagDBUrl, defaultDBUrl)
	config.JaegerURL = *loadURLParam(flagJaegerUrl, defaultJaegerUrl)
	config.SentryURL = *loadURLParam(flagSentryUrl, defaultSentryUrl)
	config.KafkaBroker = *loadURLParam(flagKafkaBroker, defaultKafkaBroker)
	config.AppID = *loadStringParam(flagSomeAppID, defaultSomeAppID)
	config.AppKey = *loadStringParam(flagSomeAppKey, defaultSomeAppID)
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует дефаултное и возращает Port
func loadPortParam(flag *int, defaultValue int) *Port {
	port := &Port{}
	// Если функция получила пустой флаг после flag.Parse() тогда возвращем порт с принятым значением по умолчанию
	if flag != nil {
		p, err := port.Create(*flag)
		// Если создание порта завершилось ошибкой, значит значение из флага не корректное, инициализируем значением
		// по-умолчанию
		if err != nil {
			log.Printf("%v, loaded default value for Port: %d \n", err, defaultValue)
			p.Value = defaultValue
		}
		// Если ошибки при создании (после валидации) нет, то возвращаем созданный Port
		return p
	}

	// Если флаг пустой тогда инициализируем Port значением по умолчанию и возвращаем
	port.Value = defaultValue
	return port
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует дефаултное и возращает URL
func loadURLParam(flag *string, defaultValue string) *URL {
	url := &URL{}
	if flag != nil {
		u, err := url.Create(*flag)
		// Если создание URL завершилось ошибкой, значит значение из флага не корректное, инициализируем значением
		// по-умолчанию
		if err != nil {
			log.Printf("%v, loaded default value for URL: %v \n", err, defaultValue)
			u.Value = defaultValue
		}
		// Если ошибки при создании (после валидации) нет, то возвращаем созданный URL
		return u
	}

	// Если флаг пустой тогда инициализируем URL значением по умолчанию и возвращаем
	url.Value = defaultValue
	return url
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает String
func loadStringParam(flag *string, defaultValue string) *string {
	if *flag != "" {
		return flag
	}
	return &defaultValue
}

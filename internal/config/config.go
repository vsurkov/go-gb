package ismgExporter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/goware/urlx"
	"github.com/namsral/flag"
)

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

// Определяем значения по умолчанию, нужны отдельно что бы к ним вернуться при не валидности полученного
const (
	defaultJSONConfig  = "configs//ismg_exporter.json"
	defaultYAMLConfig  = "configs//ismg_exporter.yaml"
	defaultPort        = 9010
	defaultDBUrl       = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	defaultJaegerUrl   = "http://jaeger:16686"
	defaultSentryUrl   = "http://sentry:9000"
	defaultKafkaBroker = "kafka:9000"
	defaultSomeAppID   = ""
	defaultSomeAppKey  = ""
)

// Определяем флаги
var (
	//flagConfigPath  = flag.String("config", defaultJsonConfig, "Configuration path with filename, example: configs/ismg_exporter.json")
	flagPort        = flag.Int("port", 0, "Server port, must be in the range 1000-65535")
	flagDBUrl       = flag.String("db_url", "", "Database connection string, example: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	flagJaegerUrl   = flag.String("jaeger_url", "", "Jaeger URL, example: https://jaeger:16686k")
	flagSentryUrl   = flag.String("sentry_url", "", "Sentry URL, example: https://sentry:9000")
	flagKafkaBroker = flag.String("kafka_broker", "", "Kafka broker URI, example: kafka:9092")
	flagSomeAppID   = flag.String("app_id", "", "Application ID")
	flagSomeAppKey  = flag.String("app_key", "", "Application KEY")
	flagDebugMode   = flag.Bool("debug_mode", false, "Debug mode enable logging, default false")
	debugMode       = false
)

// Load Парсим флаги и загружаем через функции загрузки параметры конфигурации

func (config *Config) Load(configPath string) {
	flag.Parse()

	if configPath == "" {
		configPath = defaultJSONConfig
	}

	err := config.LoadFromFile(configPath)
	if err != nil {
		log.Printf("Config from file not loaded: %v", err)
	}

	// Каждый из параметров заполняем результатом загрузки параметра
	debugMode = *flagDebugMode
	config.Port = *loadPortParam(&config.Port, flagPort, defaultPort)
	config.DbUrl = *loadURLParam(&config.DbUrl, flagDBUrl, defaultDBUrl)
	config.JaegerUrl = *loadURLParam(&config.JaegerUrl, flagJaegerUrl, defaultJaegerUrl)
	config.SentryUrl = *loadURLParam(&config.SentryUrl, flagSentryUrl, defaultSentryUrl)
	config.KafkaBroker = *loadURLParam(&config.KafkaBroker, flagKafkaBroker, defaultKafkaBroker)
	config.AppId = *loadStringParam(&config.AppId, flagSomeAppID, defaultSomeAppID)
	config.AppKey = *loadStringParam(&config.AppKey, flagSomeAppKey, defaultSomeAppKey)
}

// LoadFromFile Загружаем конфигурацию из файла
func (config *Config) LoadFromFile(configPath string) error {
	// Открываем файл по configPath и проверяем на ошибки
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// Закрываем открытый файл в любом случае
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("can't close file. %v", err)
		}
	}()

	// Декодируем конфигурацию и заполняем структуру конфигурации
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return fmt.Errorf("can't decode JSON configuration on configPath: %v. Error: \n %v", configPath, err)
	}

	return nil
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает Int
func loadPortParam(configValue *int, flag *int, defaultValue int) *int {

	if flag != nil {
		if isPortValid(flag) {
			return flag
		}
		if isPortValid(configValue) {
			return configValue
		}
		log.Printf("Port %v invalid (must be int and in the range 1000-65535), loaded default value: %d \n", *flag, defaultValue)
		return &defaultValue
	}

	if isPortValid(configValue) {
		return configValue
	}
	return &defaultValue
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает String
func loadURLParam(configValue *string, flag *string, defaultValue string) *string {
	if flag != nil {
		if isUrlValid(flag) {
			return flag
		}
		if isUrlValid(configValue) {
			return configValue
		}
		log.Printf("url invalid, loaded default value: %v \n", defaultValue)
		return &defaultValue
	}

	if isUrlValid(configValue) {
		return configValue
	}
	return &defaultValue
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует принятое значение по умолчанию
// и возращает String
func loadStringParam(configValue *string, flag *string, defaultValue string) *string {
	if *flag != "" {
		return flag
	}
	if *configValue != "" {
		return configValue
	}
	log.Printf("param invalid, loaded default value: %v \n", defaultValue)
	return &defaultValue
}

// Валидация порта, проверка, что указанное значение порта находится в разрешенном диапазоне от 1000 до 65535
func isPortValid(value *int) bool {
	if *value <= 1000 || *value > 65535 {
		return false
	}
	return true
}

// Валидация URL с учетом заданного порта
func isUrlValid(raw *string) bool {
	url, err := urlx.Parse(*raw)
	if err != nil || url.Scheme == "" || url.Host == "" {
		if debugMode {
			log.Printf("URL %v: %v", *raw, err)
		}
		return false
	}

	// Проверяем валидность порта в url
	_, portStr, _ := urlx.SplitHostPort(url)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		if debugMode {
			log.Printf("Port %v invalid (must be int and in the range 1000-65535), error: %v \n", port, err)
		}
		return false
	}

	if err != nil || !isPortValid(&port) {
		if debugMode {
			log.Printf("Port %v invalid (must be int and in the range 1000-65535) \n", port)
		}
		return false
	}
	return true
}

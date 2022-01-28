package ismgExporter

import "fmt"
import "github.com/namsral/flag"

type Config struct {
	Port                                         Port
	DB_url, Jaeger_url, Sentry_url, Kafka_broker URL
	App_id, App_key                              string
}

func NewConfig() *Config {
	return &Config{
		Port:         Port{0},
		DB_url:       URL{""},
		Jaeger_url:   URL{""},
		Sentry_url:   URL{""},
		Kafka_broker: URL{""},
		App_id:       "",
		App_key:      "",
	}
}

const (
	// Определяем значения по умолчанию, нужны отдельно что бы к ним вернуться при не валидности полученного
	defaultPort        = 9010
	defaultDBUrl       = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	defaultJaegerUrl   = "http://jaeger:16686"
	defaultSentryUrl   = "http://sentry:9000"
	defaultKafkaBroker = "kafka:9000"
	defaultSomeAppID   = "DEFAULT_APP_ID"
	defaultSomeAppKey  = "8787928792847928749248742479724"
)

var (
	// Определяем флаги
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

// Парсим флаги и загружаем через функции загрузки параметры конфигурации
func (config *Config) Load() {
	flag.Parse()

	config.Port = *loadPortParam(flagPort, defaultPort)
	config.DB_url = *loadURLParam(flagDBUrl, defaultDBUrl)
	config.Jaeger_url = *loadURLParam(flagJaegerUrl, defaultJaegerUrl)
	config.Sentry_url = *loadURLParam(flagSentryUrl, defaultSentryUrl)
	config.Kafka_broker = *loadURLParam(flagKafkaBroker, defaultKafkaBroker)
	config.App_id = *loadStringParam(flagSomeAppID, defaultSomeAppID)
	config.App_key = *loadStringParam(flagSomeAppKey, defaultSomeAppID)
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует дефаултное и возращает Port
func loadPortParam(flag *int, defaultValue int) *Port {
	port := &Port{}
	if flag != nil {
		p, err := port.Create(*flag)
		if err != nil {
			fmt.Printf("%v, loaded default value for Port: %d \n", err, defaultValue)
			p.Value = defaultValue
		}
		return p
	}
	port.Value = defaultValue
	return port
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует дефаултное и возращает URL
func loadURLParam(flag *string, defaultValue string) *URL {
	url := &URL{}
	if flag != nil {
		u, err := url.Create(*flag)
		if err != nil {
			fmt.Printf("%v, loaded default value for URL: %v \n", err, defaultValue)
			u.Value = defaultValue
		}
		return u
	}
	url.Value = defaultValue
	return url
}

// Для функция производит проверку флага и загрузку из него или из ENV, иначе использует дефаултное и возращает String
func loadStringParam(flag *string, defaultValue string) *string {
	if *flag != "" {
		return flag
	}
	return &defaultValue
}

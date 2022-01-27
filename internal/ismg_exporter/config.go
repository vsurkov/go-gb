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

// Загрузка значений из флаков или переменных, если значение не валидно то загружаем дефаултное
func (config *Config) Load() {
	var (
		// Определяем значения по умолчанию, нужны отдельно что бы к ним вернуться при не валидности полученного
		defaultPort        = 9010
		defaultDBUrl       = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
		defaultJaegerUrl   = "http://jaeger:16686"
		defaultSentryUrl   = "http://sentry:9000"
		defaultKafkaBroker = "kafka:9000"
		defaultSomeAppID   = "DEFAULT_APP_ID"
		defaultSomeAppKey  = "8787928792847928749248742479724"

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

	flag.Parse()

	if flagPort != nil {
		port := &Port{}
		port, err := port.Create(*flagPort)
		if err != nil {
			fmt.Printf("%v, loaded default value for Port \n", err)
			config.Port.Value = defaultPort
		} else {
			config.Port = *port
		}

	}

	if *flagDBUrl != "" {
		url := &URL{}
		url, err := url.Create(*flagDBUrl)
		if err != nil {
			fmt.Printf("%v, loaded default value for DB_url \n", err)
			config.DB_url.Value = defaultDBUrl
		} else {
			config.DB_url = *url
		}

	}

	if *flagJaegerUrl != "" {
		url := &URL{}
		url, err := url.Create(*flagJaegerUrl)

		if err != nil {
			fmt.Printf("%v, loaded default value for Jaeger_url \n", err)
			config.Jaeger_url.Value = defaultJaegerUrl
		} else {
			config.Jaeger_url = *url
		}

	}

	if *flagSentryUrl != "" {
		url := &URL{}
		url, err := url.Create(*flagSentryUrl)

		if err != nil {
			fmt.Printf("%v, loaded default value for Sentry_url \n", err)
			config.Sentry_url.Value = defaultSentryUrl
		} else {
			config.Sentry_url = *url
		}
	}

	if *flagKafkaBroker != "" {
		url := &URL{}
		url, err := url.Create(*flagKafkaBroker)

		if err != nil {
			fmt.Printf("%v, loaded default value for Kafka_broker \n", err)
			config.Kafka_broker.Value = defaultKafkaBroker
		} else {
			config.Kafka_broker = *url
		}
	}
	if *flagSomeAppID != "" {
		config.App_id = *flagSomeAppID
	}

	if *flagSomeAppKey != "" {
		config.App_key = *flagSomeAppKey
	}
}

//
//if *flagKafkaBroker != "" {
//	url := &URL{}
//	url, err := url.Create(*flagKafkaBroker)
//
//	if err != nil {
//		fmt.Printf("%v, loaded default value for Kafka_broker \n", err)
//		config.Kafka_broker.Value = defaultKafkaBroker
//	} else {
//		config.Kafka_broker = *url
//	}
//}
//
//func loadField(flag *int) {
//	value := flag
//	fmt.Printf(string(*value))
//}

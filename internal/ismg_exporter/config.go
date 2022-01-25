package ismgExporter

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

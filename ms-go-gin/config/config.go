package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	SupabaseUrl    string `env:"SUPABASE_URL" envDefault:""`
	SupabaseApikey string `env:"SUPABASE_API_KEY" envDefault:""`
	DsnPostgres    string `env:"DSN_POSTGRES" envDefault:""`
	Port           string `env:"PORT" envDefault:"8080"`
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   "Config.New.01",
			"error": err.Error(),
		}).Error("loading env file failed")
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   "Config.New.02",
			"error": err.Error(),
		}).Error("parsing env failed")
	}

	//logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: true})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)
	logrus.StandardLogger()
	return &cfg
}

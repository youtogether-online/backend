package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

type Config struct {
	Prod int `yaml:"prod" env:"PROD" env-default:"0"`

	Session struct {
		CookieName        string        `yaml:"cookie_name" env:"COOKIE_NAME" env-default:"session_id"`
		CookiePath        string        `yaml:"cookie_path" env:"COOKIE_PATH" env-default:"/api"`
		DurationInSeconds int           `yaml:"refresh_duration_in_seconds" env-required:"true"`
		Duration          time.Duration `yaml:"duration" env-required:"true"`
	} `yaml:"session"`

	Listen struct {
		Host string `yaml:"host" env-required:"true"`
		Port int    `yaml:"port" env-required:"true"`
	} `yaml:"listen"`

	DB struct {
		Postgres struct {
			Username string `yaml:"username" env:"POSTGRES_USERNAME" env-default:"postgres"`
			DBName   string `yaml:"db_name" env:"POSTGRES_DB" env-required:"true"`
			Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-required:"true"`
			Host     string `yaml:"host" env-required:"true"`
			Port     int    `yaml:"port" env-required:"true"`
		} `yaml:"postgres"`

		Redis struct {
			DbId int    `yaml:"db_id" env-default:"0"`
			Host string `yaml:"host" env-required:"true"`
			Port int    `yaml:"port" env-required:"true"`
		} `yaml:"redis"`
	} `yaml:"db"`

	Email struct {
		From     string `yaml:"from" env-required:"true"`
		User     string `yaml:"user" env:"EMAIL_USER" env-required:"true"`
		Password string `yaml:"password" env:"EMAIL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"EMAIL_HOST" env-default:"smtp.gmail.com"`
		Port     int    `yaml:"port" env:"EMAIL_PORT" env-default:"587"`
	} `yaml:"email"`
}

var (
	inst Config
	once sync.Once
)

// GetConfig builds the configuration file in golang type and returns it
func GetConfig() *Config {
	once.Do(func() {
		godotenv.Load()
		logrus.Println(os.Environ())

		if err := cleanenv.ReadConfig("configs/config.yml", &inst); err != nil {
			logrus.WithError(err).Error("error occurred while reading config file")
			help, _ := cleanenv.GetDescription(&inst, nil)
			logrus.Info(help)
			logrus.Exit(0)
		}

		if inst.Prod == 1 {
			inst.DB.Postgres.Host = "postgres"
			inst.DB.Redis.Host = "redis"
		}
	})

	return &inst
}

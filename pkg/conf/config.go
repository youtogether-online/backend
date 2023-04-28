package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Config struct {
	Prod int `yaml:"prod" env:"PROD" env-default:"0"`

	Session struct {
		CookieName string        `yaml:"cookie_name" env:"COOKIE_NAME" env-default:"session_id"`
		CookiePath string        `yaml:"cookie_path" env:"COOKIE_PATH" env-default:"/api"`
		Duration   time.Duration `yaml:"duration" env:"COOKIE_DURATION" env-required:"true"`
	} `yaml:"session"`

	Listen struct {
		Host string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
		Port int    `yaml:"port" env:"PORT" env-required:"true"`
	} `yaml:"listen"`

	DB struct {
		Postgres struct {
			Username string `yaml:"username" env:"POSTGRES_USERNAME" env-default:"postgres"`
			DBName   string `yaml:"db_name" env:"POSTGRES_DB" env-default:"you-together"`
			Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
			Host     string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
			Port     int    `yaml:"port" env:"POSTGRES_PORT" env-required:"true"`
		} `yaml:"postgres"`

		Redis struct {
			DbId int    `yaml:"db_id" env:"REDIS_DB" env-default:"0"`
			Host string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
			Port int    `yaml:"port" env:"REDIS_POST" env-required:"true"`
		} `yaml:"redis"`
	} `yaml:"db"`

	Email struct {
		From     string `yaml:"from" env:"EMAIL_FROM" env-required:"true"`
		User     string `yaml:"user" env:"EMAIL_USER"`
		Password string `yaml:"password" env:"EMAIL_PASSWORD"`
		Host     string `yaml:"host" env:"EMAIL_HOST"`
		Port     int    `yaml:"port" env:"EMAIL_PORT"`
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

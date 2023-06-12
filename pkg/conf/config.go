package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"sync"
	"time"
)

type Config struct {
	// in docker image is always 1
	// can be specified by environment variable
	Prod int `yaml:"prod" env:"PROD" env-default:"0"`

	Session struct {
		CookieName string        `yaml:"cookie_name" env:"COOKIE_NAME" env-default:"session_id"`
		CookiePath string        `yaml:"cookie_path" env:"COOKIE_PATH" env-default:"/api"`
		Domain     string        `yaml:"domain_name" env:"DOMAIN_NAME" env-default:"localhost"`
		Duration   time.Duration `yaml:"duration" env:"COOKIE_DURATION" env-default:"720h"`
	} `yaml:"session"`

	Listen struct {
		QueryPath string `yaml:"query_path" env:"QUERY_PATH" env-default:"/api"`
		Port      int    `yaml:"port" env:"PORT" env-default:"3000"`
	} `yaml:"listen"`

	DB struct {
		Postgres struct {
			Username string `yaml:"username" env:"POSTGRES_USERNAME" env-default:"postgres"`
			DBName   string `yaml:"db_name" env:"POSTGRES_DB" env-default:"you-together"`
			Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
			// if prod=1, host will always be "postgres" (docker constant)
			Host string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
			Port int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
		} `yaml:"postgres"`

		Redis struct {
			DbId int `yaml:"db_id" env:"REDIS_DB" env-default:"0"`
			// if prod=1, host will always be "redis" (docker constant)
			Host string `yaml:"host" env:"REDIS_HOST" env-default:"localhost"`
			Port int    `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
		} `yaml:"redis"`
	} `yaml:"db"`

	Email struct {
		From     string `yaml:"from" env:"EMAIL_FROM" env-default:"you-together@gmail.com"`
		User     string `yaml:"user" env:"EMAIL_USER"`
		Password string `yaml:"password" env:"EMAIL_PASSWORD"`
		Host     string `yaml:"host" env:"EMAIL_STMP_HOST"`
		Port     int    `yaml:"port" env:"EMAIL_PORT"`
	} `yaml:"email"`
}

var (
	inst = new(Config)
	once sync.Once
)

// GetConfig builds the golang type by environment variables
// or (if not specified) configuration file and returns it
func GetConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		if err := cleanenv.ReadConfig("configs/config.yml", inst); err != nil {
			log.WithErr(err).Err("error occurred while reading config file")
			help, _ := cleanenv.GetDescription(inst, nil)
			log.Fatal(help)
		}

		if inst.Prod == 1 {
			inst.DB.Postgres.Host = "postgres"
			inst.DB.Redis.Host = "redis"
		} else {
			log.SetLevel(log.DebugLevel)
		}
	})

	return inst
}

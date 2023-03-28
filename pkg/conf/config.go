package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Config struct {
	Prod *bool `yaml:"prod" env-required:"true"`

	Session struct {
		Secret            string        `yaml:"secret" env-required:"true"`
		CookieName        string        `yaml:"cookie_name" env-default:"session_id"`
		DurationInSeconds int           `yaml:"refresh_duration_in_seconds" env-default:"2592000"`
		Duration          time.Duration `yaml:"duration" env-default:"720h"`
	} `yaml:"session"`

	Listen struct {
		Host string `yaml:"host" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`

	DB struct {
		Postgres struct {
			Username string `yaml:"username" env-required:"true"`
			DBName   string `yaml:"db_name" env-required:"true"`
			Password string `yaml:"password" env-required:"true"`
			Host     string `yaml:"host" env-default:"127.0.0.1"`
			Port     string `yaml:"port" env-default:"5432"`
		} `yaml:"postgres"`

		Redis struct {
			Username string `yaml:"username" env-default:"redis"`
			Password string `yaml:"password" env-required:"true"`
			DB       int    `yaml:"db" env-default:"0"`
			Host     string `yaml:"host" env-default:"127.0.0.1"`
			Port     string `yaml:"port" env-default:"6379"`
		} `yaml:"redis"`
	} `yaml:"db"`

	Regexp struct {
		Email        string `yaml:"email" env-default:"."`
		Username     string `yaml:"username" env-default:"."`
		Name         string `yaml:"name" env-default:"."`
		UserPassword string `yaml:"user_password" env-default:"."`
		RoomPassword string `yaml:"room_password" env-default:"."`
		FilePath     string `yaml:"filePath" env-default:"."`
	} `yaml:"regexp"`

	Email struct {
		From     string `yaml:"from" env-required:"true"`
		User     string `yaml:"user" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		Host     string `yaml:"host" env-required:"true"`
		Port     int    `yaml:"port" env-required:"true"`
	} `yaml:"email"`
}

var inst *Config
var once sync.Once

// GetConfig builds the configuration file in golang type and returns it
func GetConfig() *Config {
	once.Do(func() {
		inst = &Config{}

		if err := cleanenv.ReadConfig("C:\\Users\\matve\\GolandProjects\\you-together\\configs\\config.yml", inst); err != nil {
			help, _ := cleanenv.GetDescription(inst, nil)
			logrus.Info(help)
			logrus.Fatalf("error occurred while reading config file: %v", err)
		}

		if inst.Regexp.Email == "." {
			logrus.Warn("email regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.Username == "." {
			logrus.Warn("username regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.Name == "." {
			logrus.Warn("names regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.UserPassword == "." {
			logrus.Warn("user password regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.RoomPassword == "." {
			logrus.Warn("room password regexp is not set. Using \".\" as a default value")
		}
		if *inst.Prod {
			inst.DB.Postgres.Host = "postgres"
			inst.DB.Redis.Host = "redis"

		}
	})

	return inst
}

package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`

	Token struct {
		Secret                   string `yaml:"secret" env-required:"true"`
		RefreshName              string `yaml:"refresh_name" env-default:"refresh_token"`
		AccessName               string `yaml:"access_name" env-default:"access_token"`
		RefreshDurationInSeconds int    `yaml:"refresh_duration_in_seconds" env-default:"604800"`
		AccessDurationInSeconds  int    `yaml:"access_duration_in_seconds" env-default:"7200"`
	} `yaml:"token"`

	Listen struct {
		Host string `yaml:"host" env-default:"127.0.0.1"`
		Port string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`

	DB struct {
		Postgres struct {
			UserName string `yaml:"user_name" env-required:"true"`
			DBName   string `yaml:"db_name" env-required:"true"`
			Password string `yaml:"password" env-required:"true"`
			Host     string `yaml:"host" env-default:"127.0.0.1"`
			Port     string `yaml:"port" env-default:"5432"`
		} `yaml:"postgres"`

		Redis struct {
			UserName string `yaml:"user_name" env-default:"redis"`
			Password string `yaml:"password" env-required:"true"`
			DB       int    `yaml:"db" env-default:"0"`
			Host     string `yaml:"host" env-default:"127.0.0.1"`
			Port     string `yaml:"port" env-default:"6379"`
		} `yaml:"redis"`
	} `yaml:"db"`

	Regexp struct {
		Email        string `yaml:"email" env-default:"."`
		UserName     string `yaml:"user_name" env-default:"."`
		Name         string `yaml:"name" env-default:"."`
		UserPassword string `yaml:"user_password" env-default:"."`
		RoomPassword string `yaml:"room_password" env-default:"."`
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

		if err := cleanenv.ReadConfig("configs/config.yml", inst); err != nil {
			help, _ := cleanenv.GetDescription(inst, nil)
			logrus.Info(help)
			logrus.Fatalf("error occurred while reading config file: %v", err)
		}

		if inst.Regexp.Email == "." {
			logrus.Warn("email regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.UserName == "." {
			logrus.Warn("user name regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.Name == "." {
			logrus.Warn("name regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.UserPassword == "." {
			logrus.Warn("user password regexp is not set. Using \".\" as a default value")
		}
		if inst.Regexp.RoomPassword == "." {
			logrus.Warn("room password regexp is not set. Using \".\" as a default value")
		}
		if inst.DB.Postgres.Host == "docker" {
			inst.DB.Postgres.Host = "postgres"
		}
		if inst.DB.Redis.Host == "docker" {
			inst.DB.Redis.Host = "redis"
		}
	})

	return inst
}

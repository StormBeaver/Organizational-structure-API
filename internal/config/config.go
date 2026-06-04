package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Database - contains all parameters database connection.
type Database struct {
	Host       string `env:"PG_HOST"`
	Port       string `env:"PG_PORT"`
	User       string `env:"PG_USER"`
	Password   string `env:"PG_PASS"`
	Migrations string `yaml:"migrations"`
	Name       string `env:"PG_NAME"`
	SslMode    string `yaml:"sslmode"`
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Port string `yaml:"port"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug bool   `yaml:"debug"`
	Name  string `yaml:"name"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project  Project `yaml:"project"`
	Rest     Rest    `yaml:"rest"`
	Database Database
}

func ReadConfig() (*Config, error) {
	env, config, err := getConfigFiles()
	if err != nil {
		return nil, err
	}

	if err := godotenv.Load(env); err != nil {
		return nil, err
	}

	var cfg Config
	if err := cleanenv.ReadConfig(config, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// ReadConfigYML - read configurations from file and init instance Config.
func getConfigFiles() (string, string, error) {
	var env, conf string

	flag.StringVar(&env, "env", ".env", "path to environment file")
	flag.StringVar(&conf, "config", "", "path to config file")
	flag.Parse()

	if conf == "" {
		return "", "", fmt.Errorf("config file not set. Use -config")
	}

	if _, err := os.Stat(env); os.IsNotExist(err) {
		return "", "", fmt.Errorf("environment file does not exist: %v", conf)
	}

	if _, err := os.Stat(conf); os.IsNotExist(err) {
		return "", "", fmt.Errorf("config file does not exist: %v", conf)
	}

	return env, conf, nil
}

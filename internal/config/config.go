package config

import (
	"os"
	"reblog/internal/log"

	"gopkg.in/yaml.v3"
)

const configFile = "reblog.yml"

var configInstance *Config

type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Prefork bool   `yaml:"prefork"`
}

type DBConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
}

type DashboardConfig struct {
	Enable bool `yaml:"enable"`
}

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	DB        DBConfig        `yaml:"db"`
	Dashboard DashboardConfig `yaml:"dashboard"`
}

func (c *Config) SaveConfig() {
	configFile, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	configBytes, err := yaml.Marshal(c)

	if err != nil {
		panic(err)
	}

	_, err = configFile.Write(configBytes)

	if err != nil {
		panic(err)
	}
}

func (c *Config) Load() error {
	configFile, err := os.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, c)

	return err
}

func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:    3000,
			Prefork: true,
		},
		DB: DBConfig{
			Type:     "mysql",
			Host:     "localhost",
			Port:     3306,
			Name:     "reblog",
			User:     "reblog",
			Password: "reblog",
		},
		Dashboard: DashboardConfig{
			Enable: true,
		},
	}
}

func NewFromFile() *Config {
	if configInstance != nil {
		return configInstance
	}

	_, err := os.Stat(configFile)

	if os.IsNotExist(err) {
		DefaultConfig().SaveConfig()
	}

	if err != nil {
		log.Panicf("[config] 配置文件加载失败: %v", err)
	}

	config := &Config{}

	config.Load()

	configInstance = config

	return config
}

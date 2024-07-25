package config

import (
	"os"
	"reblog/internal/log"

	"gopkg.in/yaml.v3"
)

const configFile = "reblog.yml"

var configInstance *Config

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	// Prefork bool   `yaml:"prefork"`
}

type DBConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	SSL      bool   `yaml:"ssl"`
}

type DashboardConfig struct {
	Enable bool `yaml:"enable"`
}

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	DB        DBConfig        `yaml:"db"`
	Dashboard DashboardConfig `yaml:"dashboard"`
	Plugins   []string        `yaml:"plugins"`
}

func (c *Config) SaveConfig() {
	configFile, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Panicf("[CONFIG] 读取配置失败: %v", err)
	}

	defer configFile.Close()

	configBytes, err := yaml.Marshal(c)

	if err != nil {
		log.Panicf("[CONFIG] 序列化配置失败: %v", err)
	}

	_, err = configFile.Write(configBytes)

	if err != nil {
		log.Panicf("[CONFIG] 配置写入失败: %v", err)
	}
}

func (c *Config) Load() error {
	configFile, err := os.ReadFile(configFile)

	configString := string(configFile)

	cookEnv(&configString)

	if err != nil {
		log.Panicf("[CONFIG] 加载配置失败%v", err)
	}

	configBytes := []byte(configString)

	err = yaml.Unmarshal(configBytes, c)

	return err
}

func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: 3000,
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
		log.Panicf("[CONFIG] 配置文件加载失败: %v", err)
	}

	config := &Config{}

	config.Load()

	configInstance = config

	return config
}

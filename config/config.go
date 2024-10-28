package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

var globalConfig *Config

type (
	Config struct {
		App        `yaml:"app"`
		Datasource `yaml:"datasource"`
	}

	App struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
		Url  string `yaml:"url"`
	}

	Datasource struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

func InitConfig() error {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("config/application.yml", cfg); err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	globalConfig = cfg
	globalConfig = cfg
	fmt.Println("***********************************************")
	fmt.Println("配置文件读取完成, 当前运行环境为: ", globalConfig.App.Env)
	fmt.Println("***********************************************")
	return nil
}

func Cfg() *Config {
	return globalConfig
}

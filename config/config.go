package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

var globalConfig *Config

type (
	Config struct {
		App        `yaml:"app"`
		Jwt        `yaml:"jwt"`
		Datasource `yaml:"datasource"`
		Session    `yaml:"session"`
		Redis      `yaml:"redis"`
	}

	App struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
		Url  string `yaml:"url"`
	}

	Jwt struct {
		Secret string `yaml:"secret"`
		Expire int64  `yaml:"expire"`
		Issuer string `yaml:"issuer"`
	}

	Session struct {
		Secret string `yaml:"secret"`
		Age    int    `yaml:"age"`
		Name   string `yaml:"name"`
	}

	Datasource struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Redis struct {
		DB       int    `yaml:"db"`
		Addr     string `yaml:"addr"`
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

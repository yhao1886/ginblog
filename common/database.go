package common

import (
	"context"
	"fmt"
	"ginblog/config"
	"ginblog/model"
	"log"
	"log/slog"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.Cfg()
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		cfg.Datasource.Username,
		cfg.Datasource.Password,
		cfg.Datasource.Host,
		cfg.Datasource.Port,
		cfg.Datasource.Database,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	DB = db

	if err := db.AutoMigrate(
		&model.Article{},
		&model.Category{},
		&model.Config{},
		&model.Menu{},
		&model.Message{},
		&model.Resource{},
		&model.Role{},
		&model.Tag{},
		&model.UserAuth{},
		&model.UserInfo{},
		&model.Page{},
		&model.Comment{},
	); err != nil {
		slog.Error(err.Error())
	}

}

func GetDB() *gorm.DB {
	return DB
}

var rdb *redis.Client

func InitRedis() {
	conf := config.Cfg()
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Redis.Password,
		DB:       conf.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Println("Redis 连接成功", conf.Redis.Addr, conf.Redis.DB, conf.Redis.Password)
}

func GetRdb() *redis.Client {
	return rdb
}

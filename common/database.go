package common

import (
	"fmt"
	"ginblog/config"
	"ginblog/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.Cfg()
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
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
	); err != nil {
		panic(err)
	}

}

func GetDB() *gorm.DB {
	return DB
}

package data

import (
	"fmt"

	"github.com/HsiaoCz/foo/User/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s??charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Mysql.User, config.Conf.Mysql.Password, config.Conf.Mysql.Host, config.Conf.Mysql.Port, config.Conf.Mysql.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("open mysql err:", zap.Error(err))
		return err
	}
	DB = db
	return nil
}

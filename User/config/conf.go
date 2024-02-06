package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(UserServiceConf)

type UserServiceConf struct {
	App   AppConf   `mapstructure:"app"`
	Mysql MysqlConf `mapstructure:"mysql"`
	Log   LogConf   `mapstructure:"log"`
	Redis RedisConf `mapstructure:"redis"`
}

type AppConf struct {
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Addr      string `mapstructure:"addr"`
	Stime     string `mapstructure:"stime"`
	MachineID int    `mapstructure:"machineID"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	Password string `mapstructure:"password"`
}

func ParseConfig() (err error) {
	viper.SetConfigName("user")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("viper read config err:", zap.Error(err))
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper unmarshal config err:", zap.Error(err))
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("config changed....")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("reunmarshal config err:", zap.Error(err))
			return
		}
	})
	return
}

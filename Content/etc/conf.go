package etc

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(FooConfig)

type FooConfig struct {
	App  AppConf `mapstructure:"app"`
	Log  LogConf `mapstructure:"log"`
	Addr Addrs   `mapstructure:"addrs"`
}

type AppConf struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	AppPort   string `mapstructure:"app_port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int    `mapstructure:"machine_id"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type Addrs struct {
	User    string `mapstructure:"user"`
	Admin   string `mapstructure:"admin"`
	Comment string `mapstructure:"comment"`
	Content string `mapstructure:"content"`
	Serach  string `mapstructure:"serach"`
}

func ParseConfig() (err error) {
	viper.SetConfigName("foo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		zap.L().Error("viper read in config err:", zap.Error(err))
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper unmarshal conf err:", zap.Error(err))
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("config changed...")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("reunmarshal config err:", zap.Error(err))
			return
		}
	})
	return
}

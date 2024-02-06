package config

type UserServiceConf struct {
	App   AppConf   `mapstructure:"app"`
	Mysql MysqlConf `mapstructure:"mysql"`
	Redis RedisConf `mapstructure:"redis"`
}

type AppConf struct {
}

type LogConf struct{}

type MysqlConf struct {
}

type RedisConf struct {
}

func ParseConfig() error {
	return nil
}

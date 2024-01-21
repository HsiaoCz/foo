package etc

import "github.com/spf13/viper"

var Conf = new(FooConfig)

type FooConfig struct {
}

func ParseConfig() {
	viper.SetConfigName("foo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../.")
}

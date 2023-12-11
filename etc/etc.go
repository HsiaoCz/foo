package etc

import "github.com/spf13/viper"

func ParseConfig(){
	viper.SetConfigName("foo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../.")
}
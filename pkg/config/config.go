package config

import "github.com/spf13/viper"

type Config struct{
	DB_URL string `mapstructure:"DB_URL"`
	PORT string `mapstructure:"PORT"`
}


func NewConfig () (config Config,err error){
	viper.AddConfigPath("./pkg/config/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	
	return
}
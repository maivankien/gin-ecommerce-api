package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		Port     int    `mapstructure:"port"`
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Port: ", config.Server.Port)

	for key, value := range config.Databases {
		fmt.Println(key, value.Host, value.Port, value.Username, value.Password)
	}
}

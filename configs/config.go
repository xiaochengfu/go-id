package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Service struct {
		Port int64 `toml:"port"`
	}
	Redis struct {
		Addr string `toml:"addr"`
		Pass string `toml:"pass"`
		Db   int    `toml:"cache"`
	}
	Cache struct {
		Driver string `toml:"driver"`
	}
}

var config = new(Config)

func init() {
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("配置读取失败： %v", err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("将配置赋给Config结构体失败：%v", err)
	}
}

func Get() Config {
	return *config
}

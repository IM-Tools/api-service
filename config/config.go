/**
  @author:panliang
  @data:2022/5/15
  @note
**/
package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConf
	Mysql  MysqlConf
	Log    LogConf
	JWT    JWTConf
}

type ServerConf struct {
	Name           string `json:"name"`
	Listen         string `json:"listen"`
	Mode           string `json:"mode"`
	Env            string `json:"env"`
	Lang           string `json:"lang"`
	CoroutinesPoll int    `json:"coroutines_poll"`
}

type JWTConf struct {
	Secret string `json:"secret"`
	Ttl    int64  `json:"ttl"`
}

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

type LogConf struct {
	Level     string `json:"level"`
	Type      string `json:"type"`
	FileName  string `json:"filename"`
	MaxSize   int    `json:"max_size"`
	MaxBackup int    `json:"max_backup"`
	MaxAge    int    `json:"max_age"`
	Compress  bool   `json:"compress"`
}

var Conf = &Config{}

// 初始化配置函数
func InitConfig(configPath string) *Config {

	// 设置文件类型
	viper.SetConfigType("yaml")

	// 读取文件配置
	viper.SetConfigFile(configPath)

	// 判断是否读取成功
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 将配置解析到Struct
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}

	// 热加载配置信息
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})

	return Conf

}

func IsLocal() bool {
	return Conf.Log.Level == "local"
}

func IsProduction() bool {
	return Conf.Log.Level == "production"
}

func IsTesting() bool {
	return Conf.Log.Level == "testing"
}

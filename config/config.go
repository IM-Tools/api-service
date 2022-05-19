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
}

type ServerConf struct {
	Listen string `json:"listen"`
	Mode   string `json:"mode"`
}

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
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

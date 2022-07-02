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
	Redis  RedisConf
	Mail   MailConf
	Kafka  KafkaConf
	Nsqe   NsqConf
}

type ServerConf struct {
	Name           string `json:"name"`
	Listen         string `json:"listen"`
	Mode           string `json:"mode"`
	Env            string `json:"env"`
	Lang           string `json:"lang"`
	CoroutinesPoll int    `json:"coroutines_poll"`
	Node           string `json:"node"`
	ServiceOpen    bool   `json:"service_open"`
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

type RedisConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	Poll     int    `json:"poll"`
	Conn     int    `json:"conn"`
}

type MailConf struct {
	Driver     string `json:"driver"`
	Host       string `json:"host"`
	Name       string `json:"name"`
	Port       string `json:"port"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
	FromName   string `json:"from_name"`
}

type KafkaConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type NsqConf struct {
	Host        string `json:"host"`
	LookupdPort string `json:"lookupd_port"`
	NsqdPort    string `json:"nsqd_port"`
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

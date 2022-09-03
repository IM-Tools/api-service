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
	Nsq    NsqConf
	QiNiu  QiNiuConfig
}

type ServerConf struct {
	Name          string `json:"name"`
	Listen        string `json:"listen"`
	Mode          string `json:"mode"`
	Env           string `json:"env"`
	Lang          string `json:"lang"`
	CoroutinePoll int    `json:"coroutinePoll"`
	Node          string `json:"node"`
	ServiceOpen   bool   `json:"serviceOpen"`
	GrpcListen    string `json:"grpcListen"`
	FilePath      string `json:"filePath"`
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
	MaxSize   int    `json:"maxSize"`
	MaxBackup int    `json:"maxBackup"`
	MaxAge    int    `json:"maxAge"`
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
	Port       int    `json:"port"`
	Password   string `json:"password"`
	Encryption string `json:"encryption"`
	FromName   string `json:"fromName"`
}

type KafkaConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type NsqConf struct {
	LookupHost string `json:"lookupHost"`
	NsqHost    string `json:"nsqHost"`
}

type QiNiuConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Domain    string `json:"domain"`
}

var Conf = &Config{}

// InitConfig 初始化配置函数
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

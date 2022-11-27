package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"im-services/internal/config"
)

func Register(address string, port int, name string, tags []string, id string) error {
	//DefaultConfig 返回客户端的默认配置
	cfg := api.DefaultConfig()

	//安装consul的ip:port
	cfg.Address = config.Conf.Consul.Host

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://10.2.69.164:5001/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func AllServices() {
	//配置
	cfg := api.DefaultConfig()
	cfg.Address = config.Conf.Consul.Host

	//将配置写入对象中
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, v := range data {
		fmt.Println("key:", key, "address:", v.Address, "port:", v.Port)
	}
}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "10.2.69.164:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//服务name
	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

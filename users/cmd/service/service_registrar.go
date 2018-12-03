package service

import (
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/log"
	"os"
	"math/rand"
	"time"
	"github.com/hashicorp/consul/api"
	consulsd "github.com/go-kit/kit/sd/consul"
	"strconv"
	)

func registerService(consulAddress string,
	consulPort string,
	advertiseAddress string,
	advertisePort string,
	logger log.Logger) (registrar sd.Registrar) {

	rand.Seed(time.Now().UTC().UnixNano())


	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = consulAddress + ":" + consulPort
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	logger.Log("advertiseAddress",advertiseAddress, "advertisePort",advertisePort)
	_http := "http://" + advertiseAddress + ":" + advertisePort + "/health"
	logger.Log("_http",_http)
	check := api.AgentServiceCheck{
		HTTP:     _http,
		Interval: "30s",
		Timeout:  "3s",
		Notes:    "Health Checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	num := rand.Intn(100) // to make service ID unique
	asr := api.AgentServiceRegistration{
		ID:      "users" + strconv.Itoa(num), //unique service ID
		Name:    "users",
		Address: advertiseAddress,
		Port:    port,
		Tags:    []string{"users", "go-microservice-base"},
		Check:   &check,
	}
	registrar = consulsd.NewRegistrar(client, &asr, logger)
	registrar.Register()
	logger.Log("users service registered")
	return registrar
}
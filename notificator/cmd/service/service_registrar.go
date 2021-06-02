package service

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
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

	logger.Log("advertiseAddress", advertiseAddress, "advertisePort", advertisePort)
	_http := "http://" + advertiseAddress + ":" + advertisePort + "/health"
	logger.Log("_http", _http)
	check := api.AgentServiceCheck{
		HTTP:     _http,
		Interval: "30s",
		Timeout:  "3s",
		Notes:    "Health Checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	num := rand.Intn(100) // to make service ID unique
	asr := api.AgentServiceRegistration{
		ID:      "notificator" + strconv.Itoa(num), //unique service ID
		Name:    "notificator",
		Address: advertiseAddress,
		Port:    port,
		Tags:    []string{"notificator", "go-microservice-base"},
		Check:   &check,
	}
	registrar = consulsd.NewRegistrar(client, &asr, logger)
	registrar.Register()
	logger.Log("notificator service registered")
	return registrar
}

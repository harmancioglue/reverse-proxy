package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

//just for testing
//can be read from yaml or any config file

var proxyConfigs = []proxyConfig{
	{
		url:    "/server1",
		server: "http://server1:8081",
	},
	{
		url:    "/server2",
		server: "http://server2:8082",
	},
	{
		url:    "/server3",
		server: "http://server3:8083",
	},
}

type proxyConfig struct {
	url    string
	server string
}

func main() {

	app := fiber.New()
	for _, config := range proxyConfigs {
		app.Use(config.url, proxy.Balancer(proxy.Config{
			Servers: []string{config.server},
		}))
	}

	app.Listen(":8087")

}

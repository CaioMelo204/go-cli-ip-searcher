package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI APP"
	app.Usage = "Get IPs and WEB servers"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Get IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "amazon.com.br",
				},
			},
			Action: getIps,
		},
		{
			Name:  "servers",
			Usage: "Get Servers",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "amazon.com.br",
				},
			},
			Action: getServers,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

//Gerar will return the command line application ready to run
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "busca de endereços na internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca o nome de servidores na internet",
			Flags:  flags,
			Action: searchServers,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

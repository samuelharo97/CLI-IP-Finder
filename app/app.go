package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the CLI app ready to be executed
func Generate() *cli.App {
	app:= cli.NewApp()
	app.Name = "IP Finder"
	app.Usage = "Searches for IP's and Server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Searches the host for it's IP addresses",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name: "servers",
			Usage: "Searches the host for it's servers",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context){
	fmt.Println("IP's are being searched for.. wait for it..")
	defer fmt.Println("All done.")

	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context){
	fmt.Println("Servers's are being searched for.. wait for it..")
	defer fmt.Println("All done.")
	
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
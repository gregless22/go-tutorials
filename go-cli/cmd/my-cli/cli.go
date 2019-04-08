package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

var (
	pl = fmt.Println
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Lets you query IPs, CNAMES, MX records"

	// flags.
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "jamesconnolly.com.au",
		},
	}

	//create the commands

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			// the action, or code that will be executed when
			// we execute our `ns` command
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("url"))
				if err != nil {
					return err
				}
				for _, item := range ns {
					pl(item.Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

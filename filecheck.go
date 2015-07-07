package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"net/http"
	"os"
	"time"
)

func checkFile(url string, warn int, crit int) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		layout := "Mon, 2 Jan 2006 15:04:05 MST"
		value := response.Header.Get("Last-Modified")
		t, _ := time.Parse(layout, value)
		timestamp := t.Unix()
		age := time.Now().Unix() - timestamp
		if int(crit) < int(age) {
			fmt.Println("Critical - File last modified on " + value)
			os.Exit(2)
		} else if int(warn) < int(age) {
			fmt.Println("Warning - File last modified on " + value)
			os.Exit(1)
		} else {
			fmt.Println("OK - File last modified on " + value)
			os.Exit(0)
		}

	}
}

func main() {
	app := cli.NewApp()
	app.Name = "http_check"
	app.Usage = "check target url"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "warn, w",
			Value: 300,
			Usage: "Specify warming value in seconds",
		},
		cli.IntFlag{
			Name:  "crit, c",
			Value: 600,
			Usage: "Specify warming value in seconds",
		},
	}

	app.Action = func(c *cli.Context) {

		if len(c.Args()) == 1 {
			url := c.Args()[0]
			checkFile(url, c.Int("warn")*60, c.Int("crit")*60)

		} else {
			fmt.Println("1 Argument required")
		}
	}

	app.Run(os.Args)

}

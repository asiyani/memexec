package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "world",
				Usage: "name for the greeting",
			},
			&cli.IntFlag{
				Name:  "sleep",
				Value: 30,
				Usage: "sleep duration in sec",
			},
			&cli.BoolFlag{
				Name:  "with_error",
				Usage: "generate error",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("hello from embeded %s\n", c.String("name"))

			fmt.Fprintf(os.Stdout, "printing on stdout\n")

			fmt.Printf("sleeping for %d\n", c.Int("sleep"))

			time.Sleep(time.Duration(c.Int("sleep")) * time.Second)

			fmt.Printf("hello from embeded after %s\n", c.String("name"))

			if c.Bool("with_error") {
				return fmt.Errorf("error generated")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

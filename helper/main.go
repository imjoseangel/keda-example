package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	helper := &cli.App{
		Name:  "helper",
		Usage: "Simple app for a keda related talk",
		Commands: []cli.Command{
			{
				Name:  "redis",
				Usage: "redis methods",
				Subcommands: []cli.Command{
					{
						Name:  "publish",
						Usage: "Publishes messages to Redis list",
						Action: func(c *cli.Context) error {
							result, err := Publish()
							if err != nil {
								fmt.Println("Failed to publish messages")
								log.Fatal(err)
							} else {
								fmt.Printf("Published messages, actual list len: %d\n", result)
							}
							return nil
						},
					},
					{
						Name:  "drain",
						Usage: "Drains the Redis list",
						Action: func(c *cli.Context) error {
							_, err := Drain()
							if err != nil {
								fmt.Println("Failed to drain Redis list")
								log.Fatal(err)
							} else {
								fmt.Println("Queue drained.")
							}
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "Prints the Redis list",
						Action: func(c *cli.Context) error {
							_, err := GetListLength()
							if err != nil {
								fmt.Println("Failed to print Redis list")
								log.Fatal(err)
							} else {
								fmt.Println("Queue printed.")
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := helper.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"os"

	"example-developer-summit/command"
	. "example-developer-summit/config"
	"example-developer-summit/kafka"
	"example-developer-summit/logger"
	"github.com/urfave/cli"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "[Demo]Developer summit <Apache Kafka>"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Yuuki Takezawa",
			Email: "takezaway@istyle.co.jp",
		},
	}
	servers := KafkaBootstrapServers
	cc := &command.CommandStruct{
		Client: kafka.KafkaClient{
			BootstrapServers: &servers,
			Logger:           &logger.ZapLogger{LogProvider: logger.Logger()},
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "produce",
			Aliases: []string{"pr"},
			Usage:   "[Demo] Apache Kafka Golang Client Producer",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, fi",
					Usage: "sample data file",
				},
			},
			Action: cc.ProduceCommand,
		},
	}
	app.Run(os.Args)
}

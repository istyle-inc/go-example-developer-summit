package command

import (
	"bufio"
	"go-example-developers-summit/kafka"
	"github.com/urfave/cli"
	"os"
)

type CommandStruct struct {
	Client kafka.KafkaClient
}

func (cmd *CommandStruct) ProduceCommand(c *cli.Context) (err error) {
	var fp *os.File
	var strs []string

	if c.String("file") == "" {
		return cli.NewExitError("Specified read file is empty", 255)
	}
	fp, err = os.Open(c.String("file"))
	defer fp.Close()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		strs = append(strs, sc.Text())
	}
	cmd.Client.PrepareProducer()
	cmd.Client.KafkaProducer.ProduceMessages(strs)
	return nil
}

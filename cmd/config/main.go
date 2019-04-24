package main

import (
	"fmt"
	"os"

	"github.com/rancher/k3os/pkg/cliinstall"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "k3os config"
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(_ *cli.Context) error {
	if os.Getuid() != 0 {
		return fmt.Errorf("must run %s as root", os.Args[0])
	}
	return cliinstall.Run()
}
package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"rizhua.com/infrastructure/injector"
)

var VERSION = "0.1.0"

func main() {
	// logger.SetVersion(VERSION)
	ctx := context.WithValue(context.Background(), "tagKey", "__main__")
	app := cli.NewApp()
	app.Name = "rizhua"
	app.Version = VERSION
	app.Usage = "RBAC scaffolding based on DDD + GIN + XORM + CASBIN + WIRE."
	app.Commands = []*cli.Command{
		newCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.WithContext(ctx).Errorf(err.Error())
	}
}

func newCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "Run web server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "server config files(.json,.yaml,.toml)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "model",
				Aliases:  []string{"m"},
				Usage:    "casbin model config(.conf)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return injector.RunServer(ctx,
				injector.SetConfigFile(c.String("conf")),
				injector.SetModelFile(c.String("model")),
				injector.SetVersion(VERSION))
		},
	}
}

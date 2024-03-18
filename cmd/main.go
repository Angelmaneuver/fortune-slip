package main

import (
	"os"

	"github.com/Angelmaneuver/fortune-slip/internal/lottery"
	"github.com/Angelmaneuver/fortune-slip/internal/web"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Fortune Slip",
		Usage: "One of the divination methods that attempts to learn about good fortune and bad fortune through divine will.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "directory",
				Aliases:  []string{"d"},
				Usage:    "`Directory Path` of the image file to be used for the fortune.",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "`Port` to be used for the http service.",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			lottery, err := lottery.New(ctx.String("d"))
			if err != nil {
				return cli.Exit(err, -1)
			}

			web.Start(ctx.Int("p"), lottery)
			return nil
		},
	}

	app.Run(os.Args)
}

package main

import (
	"fmt"
	"io"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:      "peek",
		Usage:     "peek into your shell pipelines",
		UsageText: "peek [global options]",
		// HideHelp:    true,
		HideVersion: true,
		Action: func(c *cli.Context) error {
			reader := os.Stdin
			tee := io.TeeReader(reader, os.Stdout)
			_, err := io.Copy(os.Stderr, tee)
			return err
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(1)
	}
}

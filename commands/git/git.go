package git

import (
	"fmt"

	"github.com/urfave/cli"
)

// git clone
func clone(c *cli.Context) error {
	fmt.Println("git clone invoked with args -", c.Args())
	return nil
}

// git pull
func pull(c *cli.Context) error {
	fmt.Println("git pull invoked with args -", c.Args())
	return nil
}

// New : Creates and returns cli.Command for registering it to main app
func New() cli.Command {

	return cli.Command{
		Name:    "git",
		Aliases: []string{"g"},
		Subcommands: []cli.Command{
			{
				Name:        "clone",
				Aliases:     []string{"c"},
				Usage:       "Clones the given repo",
				UsageText:   "",
				Description: "",
				Action: func(c *cli.Context) error {
					return clone(c)
				},
			},
			{
				Name:    "pull",
				Aliases: []string{"p"},
				Usage:   "Pulls the given repo",
				Action: func(c *cli.Context) error {
					return pull(c)
				},
			},
		},
	}

}

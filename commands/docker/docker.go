package docker

import (
	"fmt"

	"github.com/urfave/cli"
)

// New :
func New() cli.Command {

	return cli.Command{
		Name:    "docker",
		Aliases: []string{"d"},
		Action: func(c *cli.Context) error {
			fmt.Println("Docker command invoked !")
			return nil
		},
	}

}

package docker

import "github.com/urfave/cli"

// New :
func New() cli.Command {

	return cli.Command{
		Name: "docker",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

}

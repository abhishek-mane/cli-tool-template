package git

import "github.com/urfave/cli"

// New :
func New() cli.Command {

	return cli.Command{
		Name: "git",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

}

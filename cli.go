// cli tool template in golang
// author : Abhishek Mane <abhishekmane@outlook.in>

package main

import (
	"log"
	"os"
	"sort"

	"github.com/abhishek-mane/cli-tool-template/commands/docker"
	"github.com/abhishek-mane/cli-tool-template/commands/git"
	"github.com/abhishek-mane/cli-tool-template/config"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {

	// cli app configurations
	conf := config.Get()
	app.Name = conf["cli.name"]
	app.Usage = conf["cli.usage"]
	app.Version = conf["cli.version"]
	app.Copyright = conf["cli.copyright"]
	app.Author = conf["cli.author"]

	// global flags
	app.Flags = []cli.Flag{}

	// auto-completions
	app.EnableBashCompletion = true

	// register commands
	app.Commands = append(app.Commands, git.New())
	app.Commands = append(app.Commands, docker.New())
}

func main() {

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

var createCommand = cli.Command{
	Name:   "create",
	Usage:  "create labels from a definition file",
	Action: doCreateCommand,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "labels", Value: "default.json", Usage: "labels definition file"},
	},
}

func doCreateCommand(c *cli.Context) {
	f, err := os.Open(c.String("labels"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var labels []*github.Label
	if err := json.NewDecoder(f).Decode(&labels); err != nil {
		log.Fatal(err)
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: getGitHubToken(c)})
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	for _, repoSpec := range c.Args() {
		repo := strings.SplitN(repoSpec, "/", 2)
		for _, l := range labels {
			_, _, err := client.Issues.CreateLabel(repo[0], repo[1], l)
			if err != nil {
				log.Fatalf("Error creating label %q: %v", *l.Name, err)
			}
		}
	}
}

func getGitHubToken(c *cli.Context) string {
	if v := c.GlobalString("token"); v != "" {
		return c.GlobalString(v)
	}

	if v := c.GlobalString("token-file"); v != "" {
		if b, err := ioutil.ReadFile(v); err == nil {
			return string(b)
		}
	}

	return ""
}

func main() {
	app := cli.NewApp()
	app.Name = "label-all-the-things"
	app.Usage = "Manipulate GitHub labels"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		createCommand,
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "token",
			Usage: "GitHub API token",
		},
		cli.StringFlag{
			Name:  "token-file",
			Usage: "GitHub API token file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"os"

	endpoint "github.com/noironetworks/cilium-net/cli/endpoint"
	"github.com/noironetworks/cilium-net/cli/monitor"
	policy "github.com/noironetworks/cilium-net/cli/policy"
	"github.com/noironetworks/cilium-net/common"
	daemon "github.com/noironetworks/cilium-net/daemon"

	"github.com/codegangsta/cli"
	l "github.com/op/go-logging"
)

var (
	log = l.MustGetLogger("cilium-cli")
)

func main() {
	app := cli.NewApp()
	app.Name = "cilium"
	app.Usage = "Cilium"
	app.Version = "0.1.0"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug messages",
		},
		cli.StringFlag{
			Name:  "host, H",
			Usage: "Daemon host to connect to",
		},
	}
	app.Commands = []cli.Command{
		daemon.CliCommand,
		policy.CliCommand,
		endpoint.CliCommand,
		monitor.CliCommand,
	}
	app.Before = initEnv
	app.Run(os.Args)
}

func initEnv(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		common.SetupLOG(log, "DEBUG")
	} else {
		common.SetupLOG(log, "INFO")
	}
	return nil
}
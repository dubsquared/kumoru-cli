package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jawher/mow.cli"

	"github.com/kumoru/kumoru-cli/account"
	"github.com/kumoru/kumoru-cli/applications"
	"github.com/kumoru/kumoru-cli/configure"
	"github.com/kumoru/kumoru-cli/deployment"
	"github.com/kumoru/kumoru-cli/pools"
	"github.com/kumoru/kumoru-cli/secrets"
	"github.com/kumoru/kumoru-cli/token"
)

func init() {
	// Initialize Logging level to WARN
	// Need to change this to be configurable
	log.SetLevel(log.DebugLevel)
	///log.SetLevel(log.WarnLevel)
	log.SetOutput(os.Stderr)
}

func main() {
	app := cli.App("kumoru-cli", "Utility to interact with Kumoru services.")

	app.Command("account", "Account actions", func(act *cli.Cmd) {
		act.Command("create", "Create an account ", account.Create)
		act.Command("show", "Show acount information", account.Show)
	})

	app.Command("apps", "Application actions", func(apps *cli.Cmd) {
		apps.Command("create", "Create a new application", applications.Create)
		apps.Command("list", "List availablable applications", applications.List)
		apps.Command("show", "Show application information", applications.Show)
	})

	app.Command("configure", "Configure kumoru-cli client", configure.Cfgcli)

	app.Command("deployment", "test", deployment.Deploy)

	app.Command("pool", "Pool actions", func(pool *cli.Cmd) {
		pool.Command("create", "Create a new pool", pools.Create)
		pool.Command("list", "List availablable pools", pools.List)
		pool.Command("show", "Show pool information", pools.Show)
	})

	app.Command("secrets", "secrets actions", func(sec *cli.Cmd) {
		sec.Command("create", "Create deployments", secrets.Create)
		sec.Command("show", "Show deployment", secrets.Show)
	})

	app.Command("tokens", "Token actions", func(tkn *cli.Cmd) {
		tkn.Command("create", "Create a new token, and get a secret", token.Create)
	})

	app.Run(os.Args)
}

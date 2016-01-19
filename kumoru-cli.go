/*
Copyright 2016 Kumoru.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jawher/mow.cli"

	"github.com/kumoru/kumoru-cli/accounts"
	"github.com/kumoru/kumoru-cli/applications"
	"github.com/kumoru/kumoru-cli/pools"
	"github.com/kumoru/kumoru-cli/secrets"
	"github.com/kumoru/kumoru-cli/tokens"
)

func init() {
	// Initialize Logging level to WARN
	// Need to change this to be configurable
	log.SetLevel(log.DebugLevel)
	///log.SetLevel(log.WarnLevel)
	log.SetOutput(os.Stderr)
}

func main() {
	app := cli.App("kumoru", "Utility to interact with Kumoru services.")

	app.Version("v version", "0.0.20")

	app.Command("login", "Login action", tokens.Create)

	app.Command("accounts", "Account actions", func(act *cli.Cmd) {
		act.Command("create", "Create an account ", accounts.Create)
		act.Command("show", "Show account information", accounts.Show)
	})

	app.Command("applications", "Application actions", func(apps *cli.Cmd) {
		apps.Command("archive", "Archive an application", applications.Archive)
		apps.Command("create", "Create an application", applications.Create)
		apps.Command("deploy", "Deploy an application", applications.Deploy)
		apps.Command("list", "List all applications", applications.List)
		apps.Command("patch", "Update an application", applications.Patch)
		apps.Command("show", "Show application information", applications.Show)
	})

	app.Command("pools", "Pool actions", func(pool *cli.Cmd) {
		pool.Command("archive", "Archive a pool", pools.Archive)
		pool.Command("create", "Create a pool", pools.Create)
		pool.Command("list", "List all pools", pools.List)
		pool.Command("patch", "Update a pool", pools.Patch)
		pool.Command("show", "Show pool information", pools.Show)
	})

	app.Command("secrets", "secrets actions", func(sec *cli.Cmd) {
		sec.Command("create", "Create secret", secrets.Create)
		sec.Command("show", "Show secret information", secrets.Show)
	})

	app.Command("tokens", "Token actions", func(tkn *cli.Cmd) {
		tkn.Command("create", "Create a pair of tokens", tokens.Create)
	})

	app.Run(os.Args)
}

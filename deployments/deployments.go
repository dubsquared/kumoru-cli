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

package deployments

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/application"

	log "github.com/Sirupsen/logrus"
)

func Deploy(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Command("deploy", "Create a deployment", func(app *cli.Cmd) {
		app.Action = func() {
			resp, _, errs := application.Deploy(*uuid)
			if errs != nil {
				fmt.Println("Could not retrieve a list of applications.")
			}

			if resp.StatusCode != 202 {
				log.Fatalf("Could not deploy application: %s", resp.Status)
			}

			fmt.Sprintf("Deploying application %s", *uuid)

		}
	})
	cmd.Command("show", "Show deployment deployment information", func(app *cli.Cmd) {
		app.Action = func() {
			fmt.Println("STUB: show deployment action")
		}
	})
}

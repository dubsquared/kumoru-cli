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

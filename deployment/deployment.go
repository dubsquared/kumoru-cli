package deployment

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/service/application"
)

func Deploy(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Command("deploy", "Create deployments", func(app *cli.Cmd) {
		app.Action = func() {
			resp, body, errs := application.Deploy(*uuid)
			if errs != nil {
				fmt.Println("Could not retrieve a list of applications.")
			}

			fmt.Println(resp.StatusCode)
			utils.Pprint(body)
		}
	})
	cmd.Command("show", "Create deployments", func(app *cli.Cmd) {
		app.Action = func() {
			fmt.Println("STUB: show deployment action")
		}
	})
}

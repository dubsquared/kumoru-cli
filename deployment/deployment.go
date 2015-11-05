package deployment

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/application"
)

func Deploy(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name: "UUID",
		Desc: "Application UUID",
	})

	cmd.Command("create", "Create deployments", func(app *cli.Cmd) {
		app.Action = func() {
			application.ApplicationDeploy(*uuid)
		}
	})
	cmd.Command("show", "Create deployments", func(app *cli.Cmd) {
		app.Action = func() {
			application.ApplicationDelete(*uuid)
			fmt.Println("STUB: show deployment action")
		}
	})
	cmd.Command("delete", "Create deployments", func(app *cli.Cmd) {
		app.Action = func() {
			fmt.Println("STUB: delete deployment action")
		}
	})
}

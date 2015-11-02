package applications

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/application"
)

func List(cmd *cli.Cmd) {
	cmd.Action = func() {
		application.List()
	}
}

func Show(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("STUB: Show application information action")
	}
}

func Create(cmd *cli.Cmd) {
	name := cmd.String(cli.StringArg{
		Name: "APP_NAME",
		Desc: "Application Name",
	})

	image := cmd.String(cli.StringArg{
		Name: "IMG_URL",
		Desc: "Image URL",
	})

	enVars := cmd.Strings(cli.StringsOpt{
		Name: "e env",
		Desc: "Environment variable",
	})

	rules := cmd.Strings(cli.StringsOpt{
		Name: "r rule",
		Desc: "Application Deployment rules",
	})

	ports := cmd.Strings(cli.StringsOpt{
		Name: "p port",
		Desc: "Port",
	})

	cmd.Action = func() {
		application.Create(*name, *image, *enVars, *rules, *ports)
	}
}

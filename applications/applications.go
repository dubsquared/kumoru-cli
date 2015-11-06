package applications

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/service/application"
)

func List(cmd *cli.Cmd) {
	cmd.Action = func() {
		application.List()
	}
}

func Show(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name: "UUID",
		Desc: "Application UUID",
	})

	cmd.Action = func() {
		resp, body, errs := application.Show(*uuid)
		if errs != nil {
			fmt.Println("Could not retrieve application information.")
		}
		fmt.Println(resp.Status)

		utils.Pprint(body)
	}
}

func Create(cmd *cli.Cmd) {
	image := cmd.String(cli.StringArg{
		Name: "IMG_URL",
		Desc: "Image URL",
	})

	name := cmd.String(cli.StringArg{
		Name: "APP_NAME",
		Desc: "Application Name",
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
		resp, body, errs := application.Create(*name, *image, *enVars, *rules, *ports)
		if errs != nil {
			fmt.Println("Could not create application.")
		}
		fmt.Println(resp.Status)

		utils.Pprint(body)
	}
}

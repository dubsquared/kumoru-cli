package pools

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/service/pools"
)

func Create(cmd *cli.Cmd) {

	location := cmd.String(cli.StringOpt{
		Name:      "l location",
		Desc:      "location of pool to be created",
		EnvVar:    "KUMORU_POOL_LOCATION",
		HideValue: true,
	})

	credentials := cmd.String(cli.StringOpt{
		Name:      "c provider-credentials",
		Desc:      "Credentials of the cloud provider to be used (i.e. access_key:secret_key@aws)",
		HideValue: true,
	})

	cmd.Action = func() {
		if *location == "" || *credentials == "" {
			fmt.Println("Error: incorrect usage")
			cmd.PrintHelp()
			os.Exit(1)
		}
		resp, body, errs := pools.Create(*location, *credentials)

		if errs != nil {
			fmt.Println("Could not create a new pool.")
		}

		fmt.Println(resp.Status)

		utils.Pprint(body)
	}
}

func List(cmd *cli.Cmd) {
	cmd.Action = func() {
		resp, body, errs := pools.List()

		if errs != nil {
			fmt.Println("Could not retrieve a list of pools.")
		}

		if resp.StatusCode != 200 {
			fmt.Println(resp.Status)
		}

		utils.Pprint(body)
	}
}

func Show(cmd *cli.Cmd) {

	cmd.Action = func() {
		fmt.Println("STUB: Show pool information action")
	}

}

func Delete(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "POOL UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, body, errs := pools.Delete(*uuid)

		if errs != nil {
			fmt.Printf("Could not delete pool %s.\n", *uuid)
		}

		if resp.StatusCode != 202 {
			fmt.Println(resp.Status)
		}

		utils.Pprint(body)
	}
}

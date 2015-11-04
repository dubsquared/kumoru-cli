package pools

import (
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/pools"
)

func Create(cmd *cli.Cmd) {

	location := cmd.String(cli.StringOpt{
		Name:   "l location",
		Desc:   "location of pool to be created",
		EnvVar: "KUMORU_POOL_LOCATION",
	})

	credentials := cmd.String(cli.StringOpt{
		Name: "c provider-credentials",
		Desc: "Credentials of the cloud provider to be used (i.e. access_key:secret_key@aws)",
	})

	cmd.Action = func() {
		if *location == "" || *credentials == "" {
			fmt.Println("Error: incorrect usage")
			cmd.PrintHelp()
			os.Exit(1)
		}
		pools.Create(*location, *credentials)
	}

}

func List(cmd *cli.Cmd) {

	bootstrap := cmd.Bool(cli.BoolOpt{
		Name:  "b bootstrap",
		Value: false,
		Desc:  "Needed if you are boostraping a pool",
	})

	cmd.Action = func() {
		pools.List(*bootstrap)
	}

}

func Show(cmd *cli.Cmd) {

	cmd.Action = func() {
		fmt.Println("STUB: Show pool information action")
	}

}

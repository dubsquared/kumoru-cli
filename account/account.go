package account

import (
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
)

func Create(cmd *cli.Cmd) {

	user := cmd.String(cli.StringArg{
		Name: "USER",
		Desc: "Username",
	})

	fName := cmd.String(cli.StringOpt{
		Name: "f first-name",
		Desc: "Given Name",
	})

	lName := cmd.String(cli.StringOpt{
		Name: "l last-name",
		Desc: "Last Name",
	})

	password := cmd.String(cli.StringOpt{
		Name: "p password",
		Desc: "Password",
	})

	cmd.Action = func() {
		authorization.CreateAcct(*user, *fName, *lName, *password)
	}
}

func Show(cmd *cli.Cmd) {

	user := cmd.String(cli.StringArg{
		Name: "USER",
		Desc: "Username",
	})

	cmd.Action = func() {
		authorization.ShowAcct(*user)
	}
}

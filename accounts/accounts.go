package accounts

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
)

func Create(cmd *cli.Cmd) {

	user := cmd.String(cli.StringArg{
		Name:      "USER",
		Desc:      "Username",
		HideValue: true,
	})

	fName := cmd.String(cli.StringOpt{
		Name:      "f first-name",
		Desc:      "Given Name",
		HideValue: true,
	})

	lName := cmd.String(cli.StringOpt{
		Name:      "l last-name",
		Desc:      "Last Name",
		HideValue: true,
	})

	password := cmd.String(cli.StringOpt{
		Name:      "p password",
		Desc:      "Password",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, body, errs := authorization.CreateAcct(*user, *fName, *lName, *password)

		if errs != nil {
			fmt.Println("Could not create a new account.")
			log.Fatal(errs)
		}

		switch resp.StatusCode {
		case 200:
			fmt.Println("Account created successfully")
			fmt.Println(resp.StatusCode)
			utils.Pprint(body)
		case 409:
			fmt.Println("Account already exists.")
			fmt.Println(resp.StatusCode)
			utils.Pprint(body)
		default:
			fmt.Println(resp.StatusCode)
			utils.Pprint(body)
		}
	}
}

func Show(cmd *cli.Cmd) {

	user := cmd.String(cli.StringArg{
		Name:      "USER",
		Desc:      "Username",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, body, errs := authorization.ShowAcct(*user)
		if errs != nil {
			fmt.Println("Could not create a new account.")
		}

		switch resp.StatusCode {
		case 200:
			fmt.Println("Account created successfully")
			fmt.Println(resp.StatusCode)
			utils.Pprint(body)
		default:
			fmt.Println(resp.StatusCode)
			utils.Pprint(body)
		}

	}
}

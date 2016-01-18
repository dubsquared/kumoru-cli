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

package accounts

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
	"github.com/ryanuber/columnize"
	"golang.org/x/crypto/ssh/terminal"
)

type Account struct {
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	GivenName string `json:"given_name"`
	Surname   string `json:"surname"`
	UpdatedAt string `json:"updated_at"`
}

func Create(cmd *cli.Cmd) {
	email := cmd.String(cli.StringArg{
		Name:      "EMAIL",
		Desc:      "email address",
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

		if *password == "" {
			*password = passwordPrompt()
			fmt.Println("\n")
		}

		resp, body, errs := authorization.CreateAcct(*email, *fName, *lName, *password)

		if errs != nil {
			log.Fatalf("Could not create account: %s", errs)
		}

		switch resp.StatusCode {
		case 200:
			var account Account

			err := json.Unmarshal([]byte(body), &account)

			if err != nil {
				log.Fatal(err)
			}

			printAccountDetail(&account)
		default:
			log.Fatalf("Could not create account: %s", resp.Status)
		}
	}
}

func Show(cmd *cli.Cmd) {
	email := cmd.String(cli.StringArg{
		Name:      "EMAIL",
		Desc:      "email address",
		HideValue: true,
	})

	var account Account

	cmd.Action = func() {
		resp, body, errs := authorization.ShowAcct(*email)

		if errs != nil {
			log.Fatalf("Could not retrieve account: %s", errs)
		}

		switch resp.StatusCode {
		case 200:
			err := json.Unmarshal([]byte(body), &account)

			if err != nil {
				log.Fatal(err)
			}

			printAccountDetail(&account)
		default:
			log.Fatalf("Could not retrieve account: %s", resp.Status)
		}

	}
}

func printAccountDetail(a *Account) {
	var output []string
	fields := structs.New(a).Fields()

	fmt.Println("\nAccount Details:\n")

	for _, f := range fields {
		output = append(output, fmt.Sprintf("%s: |%s\n", f.Name(), f.Value()))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func passwordPrompt() string {
	fmt.Print("Enter password: ")
	password, errs := terminal.ReadPassword(0)

	if errs != nil {
		fmt.Println("\nCould not read password:")
		log.Fatal(errs)
		os.Exit(1)
	}

	fmt.Print("\nConfirm password: ")
	passwordConfirm, errs := terminal.ReadPassword(0)

	if errs != nil {
		fmt.Println("\nCould Not read password.")
		log.Fatal(errs)
		os.Exit(1)
	}

	if reflect.DeepEqual(password, passwordConfirm) == false {
		fmt.Println("\n")
		log.Fatal("Passwords do not match")
	}

	return strings.TrimSpace(string(password))
}

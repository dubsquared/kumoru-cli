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

package secrets

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/authorization/secrets"
	"github.com/ryanuber/columnize"
)

func Create(cmd *cli.Cmd) {
	value := cmd.String(cli.StringOpt{
		Name:      "v value",
		Desc:      "Value to be stored as secret",
		HideValue: true,
	})

	cmd.Action = func() {
		s := secrets.Secret{}

		if *value == "" {
			log.Fatal("Value must not be an empty string")
		}

		s.Value = *value
		secret, resp, errs := s.Create()

		if len(errs) > 0 {
			log.Fatalf("Could not create secret: %s", errs[0])
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Could not create secret: %s", resp.Status)
		}

		printSecretDetail(*secret)
	}
}

func Show(cmd *cli.Cmd) {
	secretUuid := cmd.String(cli.StringArg{
		Name:      "SECRET_UUID",
		Desc:      "UUID of secret to retrieve",
		HideValue: true,
	})

	cmd.Action = func() {
		s := secrets.Secret{}
		secret, resp, errs := s.Show(secretUuid)

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve secret: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve secret: %s", resp.Status)
		}

		printSecretDetail(*secret)
	}
}

func printSecretDetail(s secrets.Secret) {
	var output []string
	fields := structs.New(s).Fields()

	fmt.Println("\nSecret Details:\n")

	for _, f := range fields {
		output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), f.Value()))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

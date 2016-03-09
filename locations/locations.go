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

package locations

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/pools"
	"github.com/ryanuber/columnize"
)

func List(cmd *cli.Cmd) {
	cmd.Action = func() {
		l := pools.Location{}
		locations, resp, errs := l.List()

		if len(errs) > 0 {
			log.Fatalf("Could not retrieve locations: %s", errs[0])
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Cloud not retrieve locations: %s", resp.Status)
		}

		PrintLocationBrief(*locations)
	}
}

func PrintLocationBrief(l []pools.Location) {
	var output []string

	output = append(output, fmt.Sprintf("Location | Provider | UUID | Status"))

	for i := 0; i < len(l); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s| %s", l[i].Locate, l[i].Provider, l[i].Uuid, l[i].Status))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

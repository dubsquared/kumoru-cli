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

package regions

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/pools"
	"github.com/ryanuber/columnize"
)

type Region struct {
	CreatedAt        string `json:"created_at"`
	Credentials      string `json:"credentials"`
	Key              string `json:"key"`
	Location         string `json:"location"`
	OrchestrationUrl string `json:"orchestration_url"`
	StackId          string `json:"stack_id"`
	Status           string `json:"status"`
	UpdatedAt        string `json:"updated_at"`
	Url              string `json:"url"`
	Uuid             string `json:"uuid"`
}

func List(cmd *cli.Cmd) {
	var p []Region

	cmd.Action = func() {
		resp, body, errs := pools.List()

		if errs != nil {
			log.Fatalf("Could not retrieve regions: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Cloud not retrieve regions: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &p)

		if err != nil {
			log.Fatal(err)
		}

		PrintRegionBrief(p)
	}
}

func PrintRegionBrief(p []Region) {
	var output []string

	output = append(output, fmt.Sprintf("Location | UUID | Status"))

	for i := 0; i < len(p); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s", p[i].Location, p[i].Uuid, p[i].Status))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

package pools

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/pools"
	"github.com/ryanuber/columnize"
)

type Pool struct {
	CreatedAt        string `json:"created_at"`
	Credentials      string `json:"credentials"`
	HostLogToken     string `json:"host_log_token"`
	Key              string `json:"key"`
	Location         string `json:"location"`
	OrchestrationUrl string `json:"orchestration_url"`
	StackId          string `json:"stack_id"`
	Status           string `json:"status"`
	UpdatedAt        string `json:"updated_at"`
	Url              string `json:"url"`
	Uuid             string `json:"uuid"`
}

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
		var p Pool

		resp, body, errs := pools.Create(*location, *credentials)

		if errs != nil {
			log.Fatalf("Could not create new pool: %s", errs)
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Cloud not create new pool: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &p)

		if err != nil {
			log.Fatal(err)
		}

		printPoolBrief([]Pool{p})
		printPoolDetail(p)
	}
}

func List(cmd *cli.Cmd) {
	var p []Pool

	cmd.Action = func() {
		resp, body, errs := pools.List()

		if errs != nil {
			log.Fatalf("Could not retrieve pools: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Cloud not retrieve applications: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &p)

		if err != nil {
			log.Fatal(err)
		}

		printPoolBrief(p)
	}
}

func Show(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	var p Pool

	cmd.Action = func() {
		resp, body, errs := pools.Show(*uuid)

		if errs != nil {
			log.Fatalf("Cloud not retrieve pool: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Cloud not retrieve pool: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &p)

		if err != nil {
			log.Fatal(err)
		}

		printPoolBrief([]Pool{p})
		printPoolDetail(p)
	}
}

func Delete(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "POOL UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, _, errs := pools.Delete(*uuid)

		if errs != nil {
			log.Fatalf("Could not delete pool: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not delete pool: %s", resp.Status)
		}

		fmt.Sprintf("Pool %s accepted for archival", uuid)
	}
}

func printPoolBrief(p []Pool) {
	var output []string

	output = append(output, fmt.Sprintf("Location | Uuid | Status"))

	for i := 0; i < len(p); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s", p[i].Location, p[i].Uuid, p[i].Status))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func printPoolDetail(p Pool) {
	var output []string
	fields := structs.New(p).Fields()

	fmt.Println("\nPool Details:\n")

	for _, f := range fields {
		output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), f.Value()))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

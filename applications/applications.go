package applications

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/application"
	"github.com/ryanuber/columnize"
)

type App struct {
	Addresses           []string            `json:"addresses"`
	CreatedAt           string              `json:"created_at"`
	CurrentDeployments  map[string]string   `json:"current_deployments"`
	Environment         map[string]string   `json:"environment"`
	Hash                string              `json:"hash"`
	ImageUrl            string              `json:"image_url"`
	Location            string              `json:"pool_location"`
	LogToken            string              `json:"log_token"`
	Metadata            map[string][]string `json:"metadata"`
	OrchestrationUrl    string              `json:"orchestration_url"`
	Name                string              `json:"name"`
	PoolUuid            string              `json:"pool_uuid"`
	Ports               []string            `json:"ports"`
	ProviderCredentials string              `json:"provider_credentials"`
	Rules               map[string]int      `json:"rules"`
	Status              string              `json:"status"`
	UpdatedAt           string              `json:"updated_at"`
	Url                 string              `json:"url"`
	Uuid                string              `json:"uuid"`
	Certificates        map[string]string   `json:"certificates"`
}

type Certificates struct {
	Certificate      string `json:"certificate,omitempty"`
	PrivateKey       string `json:"private_key,omitempty"`
	CertificateChain string `json:"certificate_chain,omitempty"`
}

func Delete(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, _, errs := application.Delete(*uuid)

		if errs != nil {
			log.Fatalf("Could not delete applications: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not delete applications: %s", resp.Status)
		}

		fmt.Printf("Application %s accepted for archival\n", *uuid)
	}
}

func List(cmd *cli.Cmd) {
	var a []App

	cmd.Action = func() {
		resp, body, errs := application.List()

		if errs != nil {
			log.Fatalf("Could not retrieve applications: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve applications: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &a)

		if err != nil {
			log.Fatal(err)
		}

		printAppBrief(a)
	}
}

func Show(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	var a App

	cmd.Action = func() {
		resp, body, errs := application.Show(*uuid)

		if errs != nil {
			log.Fatalf("Could not retrieve application: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not retrieve application: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &a)

		if err != nil {
			log.Fatal(err)
		}

		printAppBrief([]App{a})
		printAppDetail(a)
	}
}

func Deploy(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, _, errs := application.Deploy(*uuid)

		if errs != nil {
			log.Fatalf("Could not deploy applications: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not deploy applications: %s", resp.Status)
		}

		fmt.Printf("Deploying application %s\n", *uuid)
	}

}

func Create(cmd *cli.Cmd) {
	poolUuid := cmd.String(cli.StringArg{
		Name: "POOL_UUID",
		Desc: "UUID of pool to create application in",
	})

	image := cmd.String(cli.StringArg{
		Name:      "IMG_URL",
		Desc:      "Image URL",
		HideValue: true,
	})

	name := cmd.String(cli.StringArg{
		Name:      "APP_NAME",
		Desc:      "Application Name",
		HideValue: true,
	})

	certificate := cmd.String(cli.StringOpt{
		Name:      "certificate_file",
		Desc:      "File(PEM encoded) containing the SSL certificate associated with the application",
		HideValue: true,
	})

	privateKey := cmd.String(cli.StringOpt{
		Name:      "private_key_file",
		Desc:      "File(PEM encoded) containing the SSL key associated with the public certificate (required if providing a certificate)",
		HideValue: true,
	})

	certificateChain := cmd.String(cli.StringOpt{
		Name:      "certificate_chain_file",
		Desc:      "File(PEM encoded) contianing the certificate chain associated with the public certificate (optional)",
		HideValue: true,
	})

	providerCredentials := cmd.String(cli.StringOpt{
		Name:      "c provider_credentials",
		Desc:      "Credentials to be used for management of application specific cloud resources (i.e. LoadBalancer, etc)",
		HideValue: true,
	})

	enVars := cmd.Strings(cli.StringsOpt{
		Name:      "e env",
		Desc:      "Environment variable",
		HideValue: true,
	})

	rules := cmd.Strings(cli.StringsOpt{
		Name:      "r rule",
		Desc:      "Application Deployment rules",
		HideValue: true,
	})

	ports := cmd.Strings(cli.StringsOpt{
		Name:      "p port",
		Desc:      "Port",
		HideValue: true,
	})

	tags := cmd.Strings(cli.StringsOpt{
		Name:      "t tags",
		Desc:      "Tags associated with the aplication being created",
		HideValue: true,
	})

	meta := cmd.Strings(cli.StringsOpt{
		Name:      "m metadata",
		Desc:      "Metadata associated with the application being created (i.e. location=cloud)",
		HideValue: true,
	})

	file := cmd.String(cli.StringOpt{
		Name:      "f file",
		Desc:      "Environment variables file",
		HideValue: true,
	})

	cmd.Action = func() {
		var a App
		var eVars []string

		fmt.Println(*file)

		if *file != "" {
			eVars = readEnvFile(*file)
		} else {
			eVars = *enVars
		}

		mData := metaData(*meta, *tags)

		certificates := readCertificates(certificate, privateKey, certificateChain)

		resp, body, errs := application.Create(*poolUuid, certificates, *name, *image, *providerCredentials, mData, eVars, *rules, *ports)
		if errs != nil {
			log.Fatalf("Could not create application: %s", errs)
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Could not create application: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &a)

		if err != nil {
			log.Fatal(err)
		}

		printAppBrief([]App{a})
		printAppDetail(a)
	}
}

func Patch(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	image := cmd.String(cli.StringOpt{
		Name:      "IMG_URL",
		Desc:      "Image URL",
		HideValue: true,
	})

	name := cmd.String(cli.StringOpt{
		Name:      "APP_NAME",
		Desc:      "Application Name",
		HideValue: true,
	})

	certificate := cmd.String(cli.StringOpt{
		Name:      "certificate_file",
		Desc:      "File(PEM encoded) containing the SSL certificate associated with the application",
		HideValue: true,
	})

	privateKey := cmd.String(cli.StringOpt{
		Name:      "private_key_file",
		Desc:      "File(PEM encoded) containing the SSL key associated with the public certificate (required if providing a certificate)",
		HideValue: true,
	})

	certificateChain := cmd.String(cli.StringOpt{
		Name:      "certificate_chain_file",
		Desc:      "File(PEM encoded) contianing the certificate chain associated with the public certificate (optional)",
		HideValue: true,
	})

	providerCredentials := cmd.String(cli.StringOpt{
		Name:      "c provider_credentials",
		Desc:      "Credentials to be used for management of application specific cloud resources (i.e. LoadBalancer, etc)",
		HideValue: true,
	})

	enVars := cmd.Strings(cli.StringsOpt{
		Name:      "e env",
		Desc:      "Environment variable",
		HideValue: true,
	})

	rules := cmd.Strings(cli.StringsOpt{
		Name:      "r rule",
		Desc:      "Application Deployment rules",
		HideValue: true,
	})

	ports := cmd.Strings(cli.StringsOpt{
		Name:      "p port",
		Desc:      "Port",
		HideValue: true,
	})

	tags := cmd.Strings(cli.StringsOpt{
		Name:      "t tags",
		Desc:      "Tags associated with the aplication being created",
		HideValue: true,
	})

	meta := cmd.Strings(cli.StringsOpt{
		Name:      "m metadata",
		Desc:      "Metadata associated with the application being created (i.e. location=cloud)",
		HideValue: true,
	})

	file := cmd.String(cli.StringOpt{
		Name:      "f file",
		Desc:      "Environment variables file",
		HideValue: true,
	})

	cmd.Action = func() {
		var a App
		var eVars []string

		fmt.Println(*file)

		if *file != "" {
			eVars = readEnvFile(*file)
		} else {
			eVars = *enVars
		}

		mData := metaData(*meta, *tags)

		certificates := readCertificates(certificate, privateKey, certificateChain)

		resp, body, errs := application.Patch(*uuid, certificates, *name, *image, *providerCredentials, mData, eVars, *rules, *ports)
		if errs != nil {
			log.Fatalf("Could not patch application: %s", errs)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Could not patch application: %s", resp.Status)
		}

		err := json.Unmarshal([]byte(body), &a)

		if err != nil {
			log.Fatal(err)
		}

		printAppDetail(a)
	}
}

func fmtRules(rules map[string]int) string {
	var r string

	for k, v := range rules {
		r = fmt.Sprintf("%s=%v ", k, v)
	}

	return r
}

func metaData(meta, tags []string) string {
	var mdata string

	if len(meta) > 0 {

		for _, data := range meta {
			e := strings.Split(data, "=")
			mdata += fmt.Sprintf("\"%s\":\"%s\",", e[0], e[1])
		}
	}

	if len(tags) > 0 {
		t, _ := json.Marshal(tags)
		mdata += fmt.Sprintf("\"tags\": %s", t)
	}

	if mdata == "" {
		return ""
	}

	return fmt.Sprintf("{%s}\n", mdata)

}

func printAppBrief(a []App) {
	var output []string

	output = append(output, fmt.Sprintf("Name | Uuid | Status | Location | Ports | Rules"))

	for i := 0; i < len(a); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s | %s | %s | %s", a[i].Name, a[i].Uuid, a[i].Status, a[i].Location, a[i].Ports, fmtRules(a[i].Rules)))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func printAppDetail(a App) {
	var output []string
	fields := structs.New(a).Fields()

	fmt.Println("\nApplication Details:\n")

	for _, f := range fields {
		if f.Name() == "Rules" {
			output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), fmtRules(a.Rules)))
		} else {
			output = append(output, fmt.Sprintf("%s: |%v\n", f.Name(), f.Value()))
		}
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func readCertificates(certificate, privateKey, certificateChain *string) string {
	var certificates Certificates

	if *certificate != "" {
		cert, err := ioutil.ReadFile(*certificate)
		if err != nil {
			log.Fatal(err)
		}
		certificates.Certificate = string(cert)
	}

	if *privateKey != "" {
		key, err := ioutil.ReadFile(*privateKey)
		if err != nil {
			log.Fatal(err)
		}
		certificates.PrivateKey = string(key)
	}

	if *certificateChain != "" {
		chain, err := ioutil.ReadFile(*certificateChain)
		if err != nil {
			log.Fatal(err)
		}
		certificates.CertificateChain = string(chain)
	}

	c, err := json.Marshal(certificates)

	if err != nil {
		log.Fatal(err)
	}

	if string(c) == "{}" {
		c = []byte("")
	}

	return string(c)
}

func readEnvFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	x := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x = append(x, scanner.Text())
	}

	fmt.Println(x)

	return x
}

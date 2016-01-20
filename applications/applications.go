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

package applications

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/application"
	"github.com/ryanuber/columnize"
)

type App struct {
	Addresses           []string               `json:"addresses"`
	CreatedAt           string                 `json:"created_at"`
	CurrentDeployments  map[string]string      `json:"current_deployments"`
	Environment         map[string]string      `json:"environment"`
	Hash                string                 `json:"hash"`
	ImageUrl            string                 `json:"image_url"`
	Location            string                 `json:"pool_location"`
	Metadata            map[string]interface{} `json:"metadata"`
	OrchestrationUrl    string                 `json:"orchestration_url"`
	Name                string                 `json:"name"`
	PoolUuid            string                 `json:"pool_uuid"`
	Ports               []string               `json:"ports"`
	ProviderCredentials string                 `json:"provider_credentials"`
	Rules               map[string]int         `json:"rules"`
	SSLPorts            []string               `json:"ssl_ports"`
	Status              string                 `json:"status"`
	UpdatedAt           string                 `json:"updated_at"`
	Url                 string                 `json:"url"`
	Uuid                string                 `json:"uuid"`
	Certificates        map[string]string      `json:"certificates"`
}

type Certificates struct {
	Certificate      string `json:"certificate,omitempty"`
	PrivateKey       string `json:"private_key,omitempty"`
	CertificateChain string `json:"certificate_chain,omitempty"`
}

func Archive(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, _, errs := application.Delete(*uuid)

		if errs != nil {
			log.Fatalf("Could not archive applications: %s", errs)
		}

		if resp.StatusCode != 202 {
			log.Fatalf("Could not archive applications: %s", resp.Status)
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
		Name:      "POOL_UUID",
		Desc:      "UUID of pool to create application in",
		HideValue: true,
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

	certificateChain := cmd.String(cli.StringOpt{
		Name:      "certificate_chain_file",
		Desc:      "File(PEM encoded) contianing the certificate chain associated with the public certificate (optional)",
		HideValue: true,
	})

	envFile := cmd.String(cli.StringOpt{
		Name:      "env_file",
		Desc:      "Environment variables file",
		HideValue: true,
	})

	privateKey := cmd.String(cli.StringOpt{
		Name:      "private_key_file",
		Desc:      "File(PEM encoded) containing the SSL key associated with the public certificate (required if providing a certificate)",
		HideValue: true,
	})

	sslPorts := cmd.Strings(cli.StringsOpt{
		Name:      "ssl_port",
		Desc:      "Port to be associated with the certificate",
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
		Desc:      "Port (non-ssl)",
		HideValue: true,
	})

	labels := cmd.Strings(cli.StringsOpt{
		Name:      "l label",
		Desc:      "Label associated with the application",
		HideValue: true,
	})

	meta := cmd.String(cli.StringOpt{
		Name:      "m metadata",
		Desc:      "Metadata associated with the application being created (i.e. location=cloud)",
		HideValue: true,
	})

	cmd.Action = func() {
		var a App
		var eVars []string

		if *envFile != "" {
			eVars = readEnvFile(*envFile)
		} else {
			eVars = *enVars
		}

		mData := metaData(*meta, *labels)

		certificates := readCertificates(certificate, privateKey, certificateChain)

		resp, body, errs := application.Create(*poolUuid, certificates, *name, *image, *providerCredentials, mData, eVars, *rules, *ports, *sslPorts)
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
		Desc:      "File (PEM) containing the SSL certificate associated with the application",
		HideValue: true,
	})

	envFile := cmd.String(cli.StringOpt{
		Name:      "env_file",
		Desc:      "Environment variables file",
		HideValue: true,
	})

	certificateChain := cmd.String(cli.StringOpt{
		Name:      "certificate_chain_file",
		Desc:      "File (PEM) contianing the certificate chain associated with the public certificate (optional)",
		HideValue: true,
	})

	privateKey := cmd.String(cli.StringOpt{
		Name:      "private_key_file",
		Desc:      "File (PEM) containing the SSL key associated with the public certificate (required if providing a certificate)",
		HideValue: true,
	})

	sslPorts := cmd.Strings(cli.StringsOpt{
		Name:      "ssl_port",
		Desc:      "Port to be assocaited with the certificate",
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

	labels := cmd.Strings(cli.StringsOpt{
		Name:      "l label",
		Desc:      "Label associated with the aplication",
		HideValue: true,
	})

	meta := cmd.String(cli.StringOpt{
		Name:      "m metadata",
		Desc:      "Metadata associated with the application being created (i.e. location=cloud)",
		HideValue: true,
	})

	cmd.Action = func() {
		var a App
		var eVars []string

		if *envFile != "" {
			eVars = readEnvFile(*envFile)
		} else {
			eVars = *enVars
		}

		mData := metaData(*meta, *labels)

		certificates := readCertificates(certificate, privateKey, certificateChain)

		resp, body, errs := application.Patch(*uuid, certificates, *name, *image, *providerCredentials, mData, eVars, *rules, *ports, *sslPorts)
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
		r += fmt.Sprintf("%s=%v ", k, v)
	}

	return r
}

//metaData combines the provided list of labels with provided arbitary metadata and asserts the result is proper JSON
//It returns the metadata JSON string
func metaData(meta string, labels []string) string {
	js := map[string]interface{}{
		"labels": []string{},
	}

	if len(meta) > 0 {
		err := json.Unmarshal([]byte(meta), &js)
		if err != nil {
			fmt.Println("metadata must be valid JSON:")
			log.Fatal(err)
		}
	}

	if len(labels) > 0 {
		js["labels"] = labels
	}

	mdata, err := json.Marshal(js)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", mdata)
}

func printAppBrief(a []App) {
	var output []string

	output = append(output, fmt.Sprintf("Name | Uuid | Status | Location | Ports | SSL Ports | Rules"))

	for i := 0; i < len(a); i++ {
		output = append(output, fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s", a[i].Name, a[i].Uuid, a[i].Status, a[i].Location, a[i].Ports, a[i].SSLPorts, fmtRules(a[i].Rules)))
	}

	fmt.Println(columnize.SimpleFormat(output))
}

func printAppDetail(a App) {
	var output []string
	fields := structs.New(a).Fields()

	fmt.Println("\nApplication Details:\n")

	for _, f := range fields {
		if f.Name() == "CurrentDeployments" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for k, v := range a.CurrentDeployments {
				output = append(output, fmt.Sprintf("……|%s: %s", k, v))
			}
		} else if f.Name() == "Metadata" {
			mdata, _ := json.Marshal(a.Metadata)
			output = append(output, fmt.Sprintf("%s: |%s\n", f.Name(), mdata))
		} else if f.Name() == "Rules" {
			output = append(output, fmt.Sprintf("%s:\n", f.Name()))
			for k, v := range a.Rules {
				output = append(output, fmt.Sprintf("……|%s=%v", k, v))
			}
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
		log.Fatal(err)
	}

	defer f.Close()

	x := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x = append(x, scanner.Text())
	}

	return x
}

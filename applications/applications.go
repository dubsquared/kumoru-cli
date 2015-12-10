package applications

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru/utils"
	"github.com/kumoru/kumoru-sdk-go/service/application"
)

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
		resp, body, errs := application.Delete(*uuid)
		if errs != nil {
			fmt.Println("Could not delete application.")
		}

		fmt.Println(resp.StatusCode)
		utils.Pprint(body)
	}
}

func List(cmd *cli.Cmd) {
	cmd.Action = func() {
		resp, body, errs := application.List()
		if errs != nil {
			fmt.Println("Could not retrieve application information.")
		}

		fmt.Println(resp.StatusCode)
		utils.Pprint(body)
	}
}

func Show(cmd *cli.Cmd) {
	uuid := cmd.String(cli.StringArg{
		Name:      "UUID",
		Desc:      "Application UUID",
		HideValue: true,
	})

	cmd.Action = func() {
		resp, body, errs := application.Show(*uuid)
		if errs != nil {
			fmt.Println("Could not retrieve application information.")
		}

		fmt.Println(resp.StatusCode)
		utils.Pprint(body)
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

		var eVars []string

		fmt.Println(*file)

		if *file != "" {
			eVars = readEnvFile(*file)
		} else {
			eVars = *enVars
		}

		mData := metaData(*meta, *tags)

		credentials := readCredentials(certificate, privateKey, certificateChain)

		resp, body, errs := application.Create(*poolUuid, credentials, *name, *image, *providerCredentials, mData, eVars, *rules, *ports)
		if errs != nil {
			fmt.Println("Could not create application.")
		}

		fmt.Println(resp.StatusCode)
		utils.Pprint(body)
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
		var eVars []string

		fmt.Println(*file)

		if *file != "" {
			eVars = readEnvFile(*file)
		} else {
			eVars = *enVars
		}

		mData := metaData(*meta, *tags)

		credentials := readCredentials(certificate, privateKey, certificateChain)

		resp, body, errs := application.Patch(*uuid, credentials, *name, *image, *providerCredentials, mData, eVars, *rules, *ports)
		if errs != nil {
			fmt.Println("Could not patch application.")
		}

		fmt.Println(resp.StatusCode)
		utils.Pprint(body)
	}
}

func readCredentials(certificate, privateKey, certificateChain *string) string {
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

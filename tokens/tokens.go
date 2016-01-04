package tokens

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
)

func Create(cmd *cli.Cmd) {
	force := cmd.Bool(cli.BoolOpt{
		Name:      "force",
		Desc:      "If a set of tokens already exist, overwrite them.",
		Value:     false,
		HideValue: true,
	})
	save := cmd.Bool(cli.BoolOpt{
		Name:      "s save",
		Desc:      "Save tokens to file",
		Value:     false,
		HideValue: true,
	})

	cmd.Action = func() {
		usrHome := os.Getenv("HOME")
		directory := usrHome + "/.kumoru/"
		filename := "config"
		kfile := directory + filename

		if kumoru.HasTokens(kfile, "tokens") == true && *force == false {
			fmt.Println(kfile, "configuration file already exists, or tokens already exist.")
			fmt.Println("Please see help for additonal options.")
			os.Exit(1)
		}

		token, resp, body, errs := authorization.GetTokens(credentials())

		if errs != nil {
			fmt.Println("Could not retrieve new tokens")
			log.Fatal(errs)
		}

		if resp.StatusCode != 201 {
			fmt.Println(fmt.Sprintf("Could not retrieve token, server responded: %v", resp.Status))
			os.Exit(1)
		}

		switch *save {

		default:
			fmt.Printf("\n[tokens]\n")
			fmt.Printf("kumoru_token_public=%s\n", token)
			fmt.Printf("kumoru_token_private=%s\n", body)

		case true:
			errs := kumoru.SaveTokens(directory, filename, "tokens", kumoru.Ktokens{
				Public:  token,
				Private: body,
			})

			if errs != nil {
				fmt.Println("Could not save tokens to file")
				log.Fatal(errs)
			}
			fmt.Printf("\nTokens saved to %s.\n", kfile)

		}
	}

}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\nGenerating new token.\n")

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, errs := terminal.ReadPassword(0)
	if errs != nil {
		fmt.Println("\nCould Not read password.")
		log.Fatal(errs)
		os.Exit(1)
	}

	fmt.Println("\nPlease wait while we generate new tokens.\n")

	return strings.TrimSpace(username), strings.TrimSpace(string(bytePassword))
}

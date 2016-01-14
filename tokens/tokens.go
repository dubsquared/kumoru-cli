package tokens

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/kumoru"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
	"golang.org/x/crypto/ssh/terminal"
)

func Create(cmd *cli.Cmd) {
	force := cmd.Bool(cli.BoolOpt{
		Name:      "f force",
		Desc:      "If a set of tokens already exist in ~/.kumoru/config, overwrite them.",
		Value:     false,
		HideValue: true,
	})
	dontSave := cmd.Bool(cli.BoolOpt{
		Name:      "d dont_save",
		Desc:      "Output tokens to screen instead of saving to configuration",
		Value:     false,
		HideValue: true,
	})

	cmd.Action = func() {
		usrHome := os.Getenv("HOME")
		directory := usrHome + "/.kumoru/"
		filename := "config"
		kfile := directory + filename

		if *dontSave == false {
			if kumoru.HasTokens(kfile, "tokens") == true && *force == false {
				fmt.Println(kfile, "configuration file already exists, or tokens already exist.")
				fmt.Println("Please see help for additonal options.")
				os.Exit(1)
			}
		}

		token, resp, body, errs := authorization.GetTokens(credentials())

		if errs != nil {
			log.Fatalf("Could not retrieve new tokens: %s", errs)
		}

		if resp.StatusCode != 201 {
			log.Fatalf("Could not retrieve tokens: %s", resp.Status)
			os.Exit(1)
		}

		switch *dontSave {
		default:
			errs := kumoru.SaveTokens(directory, filename, "tokens", kumoru.Ktokens{
				Public:  token,
				Private: body,
			})

			if errs != nil {
				log.Fatalf("Could not save tokens to file: %s", errs)
			}

			fmt.Printf("\nTokens saved to %s.\n", kfile)
		case true:
			fmt.Printf("\n[tokens]\n")
			fmt.Printf("kumoru_token_public=%s\n", token)
			fmt.Printf("kumoru_token_private=%s\n", body)
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

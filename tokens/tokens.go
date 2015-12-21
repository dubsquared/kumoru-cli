package tokens

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		file := usrHome + "/.kumoru/config"

		if kumoru.HasTokens(file, "tokens") == true && *force == false {
			fmt.Println(file, "configuration file already exists, or tokens already exist.")
			fmt.Println("Please see help for additonal options.")
			os.Exit(1)
		}

		token, resp, body, errs := authorization.GetTokens(credentials())

		if errs != nil {
			fmt.Println("Could not retrieve new tokens")
		}

		if resp.StatusCode != 201 {
			fmt.Println(fmt.Sprintf("Could not retrieve token, server responded: %v", resp.Status))
		}

		fmt.Printf("\n[tokens]\n")
		fmt.Printf("kumoru_token_public=%s\n", token)
		fmt.Printf("kumoru_token_private=%s\n", body)

		if *save {
			err := kumoru.SaveTokens(file, "tokens", kumoru.Ktokens{
				Public:  token,
				Private: body,
			})
			if err != nil {
				fmt.Println("Could not save tokens to file")
			}
		}
	}

}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\nGenerating new token.\n")

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("\nCould Not read password.")
		os.Exit(1)
	}

	fmt.Println("Please wait while we generate new tokesn.\n")

	return strings.TrimSpace(username), strings.TrimSpace(string(bytePassword))
}

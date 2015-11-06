package token

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
)

func Create(cmd *cli.Cmd) {
	cmd.Action = func() {
		token, resp, body, errs := authorization.GetTokens()

		if errs != nil {
			fmt.Println("Could not retrieve new tokens")
		}

		if resp.StatusCode != 201 {
			fmt.Println(fmt.Sprintf("Could not retrieve token, server responded: %v", resp.Status))
		}

		fmt.Printf("\n[tokens]\n")
		fmt.Printf("kumoru_token_public=%s\n", token)
		fmt.Printf("kumoru_token_private=%s\n", body)
	}

}

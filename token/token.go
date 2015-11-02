package token

import (
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/kumoru/kumoru-sdk-go/service/authorization"
)

func Create(cmd *cli.Cmd) {

	cmd.Action = func() {
		token, secret := authorization.GetTokens()
		fmt.Printf("\n[tokens]\n")
		fmt.Printf("kumoru_token_public=%s\n", token)
		fmt.Printf("kumoru_token_private=%s\n", secret)
	}

}

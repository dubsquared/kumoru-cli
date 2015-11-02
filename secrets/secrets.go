package secrets

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

func Create(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("STUB: create secrets action")
	}
}

func Show(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("STUB: Show secrets action")
	}
}

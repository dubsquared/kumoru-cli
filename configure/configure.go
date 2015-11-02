package configure

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

func Cfgcli(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println("STUB: Authorization get Tokens action")
	}

}

package configure

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"gopkg.in/ini.v1"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/jawher/mow.cli"
)

type config struct {
	user string
	pass string
	auth string
	app  string
	pool string
}

func Cfgcli(cmd *cli.Cmd) {
	cmd.Action = func() {
		usr, err := user.Current()
		validate(err)

		file := usr.HomeDir + "/.kumoru/config"
		if _, err := os.Stat(file); err == nil {
			fmt.Println(file, "configuration file already exists.")
			os.Exit(1)
		}

		info := config{}

		info.user, info.pass = credentials()
		info.auth, info.pool, info.app = endpoints()

		writeConfig(info)
	}
}

func writeConfig(conf config) {
	usr, err := user.Current()
	validate(err)

	path := usr.HomeDir + "/.kumoru/"
	filename := path + "config"

	cfg := ini.Empty()

	_, err = cfg.Section("auth").NewKey("kumoru_username", conf.user)
	validate(err)
	_, err = cfg.Section("auth").NewKey("kumoru_password", conf.pass)
	validate(err)

	_, err = cfg.Section("endpoints").NewKey("kumoru_pool_api", conf.pool)
	validate(err)
	_, err = cfg.Section("endpoints").NewKey("kumoru_application_api", conf.app)
	validate(err)
	_, err = cfg.Section("endpoints").NewKey("kumoru_authorization_api", conf.auth)
	validate(err)

	mkdir, err := exists(path)
	validate(err)

	if !mkdir {
		err = os.Mkdir(path, 0700)
		validate(err)
	}

	err = cfg.SaveTo(filename)
	validate(err)

	fmt.Println("Configuration file has been saved to: %s", filename)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func validate(err error) {

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("\nCould Not read password.")
		os.Exit(1)
	}

	return strings.TrimSpace(username), strings.TrimSpace(string(bytePassword))
}

func endpoints() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter Authorization endpoint [http://authorization.api.kumoru.io:5000]:")
	auth, _ := reader.ReadString('\n')
	if auth == "\n" {
		auth = "http://authorization.api.kumoru.io:5000"
	}

	fmt.Print("Enter Application endpoint [http://application.api.kumoru.io]:")
	apps, _ := reader.ReadString('\n')
	if apps == "\n" {
		apps = "http://authorization.api.kumoru.io"
	}

	fmt.Print("Enter Pool endpoint [http://authorization.api.kumoru.io:5000]:")
	pool, _ := reader.ReadString('\n')
	if pool == "\n" {
		pool = "http://authorization.api.kumoru.io:5000"
	}

	return strings.TrimSpace(auth), strings.TrimSpace(pool), strings.TrimSpace(apps)
}

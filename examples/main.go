package main

import (
	"fmt"
	"os"
	"rivian/internal/prompts"
	"rivian/rivian"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

func main() {
	var username, password string
	if os.Getenv("RIVIAN_USERNAME") == "" {
		username, _ = prompts.Username()
	} else {
		fmt.Println("Loaded username from environment variables...")
	}
	if os.Getenv("RIVIAN_PASSWORD") == "" {
		password, _ = prompts.Password()
	} else {
		fmt.Println("Loaded password from environment variables...")
	}

	r1T := &rivian.Rivian{
		Username: username,
		Password: password,
		Logger:   logrus.New(),
	}
	r1T.Authenticate()

	fmt.Printf("%+v", r1T)
}

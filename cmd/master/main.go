package main

import (
	"github.com/mondaycodingclub/my-nginx/cmd/master/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewMasterCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"github.com/mondaycodingclub/my-nginx/cmd/agent/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewAgentCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

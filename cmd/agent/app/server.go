package app

import (
	"context"
	"fmt"
	"github.com/mondaycodingclub/my-nginx/cmd/agent/app/config"
	"github.com/mondaycodingclub/my-nginx/pkg/agent"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// NewAgentCommand creates an object of *cobra.Command with default parameters and registryOptions
func NewAgentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "my-nginx-agent",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCommand(cmd, args); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// runCommand runs the agent.
func runCommand(cmd *cobra.Command, args []string) error {
	c := &config.Config{}
	cc := c.Complete()
	stopCh := make(chan struct{})
	return Run(cc, stopCh)
}

// Run executes the scheduler based on the given configuration. It only return on error or when stopCh is closed.
func Run(cc config.CompletedConfig, stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	log.Infof("Starting Kubernetes Scheduler version %+v", "version.Get()")

	// Create the scheduler.
	agent, err := agent.New(stopCh)
	if err != nil {
		return err
	}

	// Prepare a reusable runCommand function.
	run := func(ctx context.Context) {
		agent.Run()
		<-ctx.Done()
	}

	ctx, cancel := context.WithCancel(context.TODO()) // TODO once Run() accepts a context, it should be used here
	defer cancel()

	go func() {
		select {
		case <-stopCh:
			cancel()
		case <-ctx.Done():
		}
	}()

	run(ctx)
	return fmt.Errorf("finished without leader elect")
}

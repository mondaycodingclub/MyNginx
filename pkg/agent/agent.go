package agent

import (
	"github.com/mondaycodingclub/my-nginx/api/heartbeat"
	"github.com/mondaycodingclub/my-nginx/api/nginx"
	"github.com/mondaycodingclub/my-nginx/pkg/agent/client"
	log "github.com/sirupsen/logrus"
	"time"
)

// The Agent is the interface for agent
type Agent struct {
	AgentName      string
	NginxInstances map[string]*nginx.Instance
	MasterClient   *client.MasterClient
	StopEverything <-chan struct{} // Close this to shut down the agent.
}

type agentOptions struct {
	agentName string
}

// Option configures an Agent
type Option func(options *agentOptions)

// WithName sets agentName for Scheduler, the default agentName is hostIp + randomString
func WithName(agentName string) Option {
	return func(o *agentOptions) {
		o.agentName = agentName
	}
}

var defaultSchedulerOptions = agentOptions{
	agentName: "my-nginx-1",
}

// New returns a Scheduler
func New(stopCh <-chan struct{}, opts ...func(o *agentOptions)) (*Agent, error) {
	options := defaultSchedulerOptions
	for _, opt := range opts {
		opt(&options)
	}
	return &Agent{
		AgentName:      options.agentName,
		NginxInstances: make(map[string]*nginx.Instance),
		MasterClient:   client.NewMasterClient("", 8080, ""),
		StopEverything: stopCh,
	}, nil
}

// Run begins watching and scheduling. It waits for cache to be synced, then starts a goroutine and returns immediately.
func (agent *Agent) Run() {
	go agent.heartbeatLoop()

}

func (agent *Agent) heartbeatLoop() {
	for {
		agent.heartbeat()
		time.Sleep(10 * time.Second)
	}
}

func (agent *Agent) heartbeat() {
	log.Infof("heartbeat start")
	currentNginxInstances := make([]*nginx.Instance, len(agent.NginxInstances))
	for _, nginxInstance := range agent.NginxInstances {
		currentNginxInstances = append(currentNginxInstances, nginxInstance)
	}
	request := &heartbeat.Request{
		IP:             "0.0.0.0",
		Port:           8080,
		NginxInstances: currentNginxInstances,
	}
	response, err := agent.MasterClient.SyncHeartbeat(request)
	if err != nil {
		log.Errorf("heartbeat sync failed: %v", err)
		return
	}
	latestNginxInstances := response.NginxInstances
	for _, latest := range latestNginxInstances {
		current := agent.NginxInstances[latest.ID]
		if current.Spec.Revision > latest.Spec.Revision {
			current.Spec = latest.Spec
		}
	}
	log.Infof("heartbeat finished")
}

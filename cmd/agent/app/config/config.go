package config

// Config has all the context to run an Agent
type Config struct {
	// ComponentConfig is the scheduler server's configuration object.
	//ComponentConfig kubeschedulerconfig.KubeSchedulerConfiguration
	//
	//// LoopbackClientConfig is a config for a privileged loopback connection
	//LoopbackClientConfig *restclient.Config
	//
	//InsecureServing        *apiserver.DeprecatedInsecureServingInfo // nil will disable serving on an insecure port
	//InsecureMetricsServing *apiserver.DeprecatedInsecureServingInfo // non-nil if metrics should be served independently
	//Authentication         apiserver.AuthenticationInfo
	//Authorization          apiserver.AuthorizationInfo
	//SecureServing          *apiserver.SecureServingInfo
	//
	//Client          clientset.Interface
	//InformerFactory informers.SharedInformerFactory
	//PodInformer     coreinformers.PodInformer
	//EventClient     v1beta1.EventsGetter
	//
	//MatrixConfig                   rpc.MatrixConfiguration
	//MatrixResourceViewSyncDuration time.Duration
	//MatrixClient                   rpc.MatrixClient
	//MatrixNodeInformer             matrixinformer.MatrixNodeInformer
	//
	//// TODO: Remove the following after fully migrating to the new events api.
	//CoreEventClient           v1core.EventsGetter
	//LeaderElectionBroadcaster record.EventBroadcaster
	//
	//Recorder    events.EventRecorder
	//Broadcaster events.EventBroadcaster
	//
	//// LeaderElection is optional.
	//LeaderElection *leaderelection.LeaderElectionConfig
}

type completedConfig struct {
	*Config
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	cc := completedConfig{c}

	//if c.InsecureServing != nil {
	//	c.InsecureServing.Name = "healthz"
	//}
	//if c.InsecureMetricsServing != nil {
	//	c.InsecureMetricsServing.Name = "metrics"
	//}
	//
	//apiserver.AuthorizeClientBearerToken(c.LoopbackClientConfig, &c.Authentication, &c.Authorization)

	return CompletedConfig{&cc}
}

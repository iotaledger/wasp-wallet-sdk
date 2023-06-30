package types

// Options for the client builder
type ClientOptions struct {
	// Timeout for API requests
	APITimeout *IDuration `json:"apiTimeout,omitempty" yaml:"apiTimeout,omitempty" mapstructure:"apiTimeout,omitempty"`

	// Options for the MQTT broker
	BrokerOptions *IMqttBrokerOptions `json:"brokerOptions,omitempty" yaml:"brokerOptions,omitempty" mapstructure:"brokerOptions,omitempty"`

	// If the node health status should be ignored
	IgnoreNodeHealth bool `json:"ignoreNodeHealth,omitempty" yaml:"ignoreNodeHealth,omitempty" mapstructure:"ignoreNodeHealth,omitempty"`

	// Whether the PoW should be done locally or remotely.
	LocalPow bool `json:"localPow,omitempty" yaml:"localPow,omitempty" mapstructure:"localPow,omitempty"`

	// Minimum amount of nodes required for request when quorum is enabled
	MinQuorumSize float64 `json:"minQuorumSize,omitempty" yaml:"minQuorumSize,omitempty" mapstructure:"minQuorumSize,omitempty"`

	// Data related to the used network
	NetworkInfo *INetworkInfo `json:"networkInfo,omitempty" yaml:"networkInfo,omitempty" mapstructure:"networkInfo,omitempty"`

	// Interval in which nodes will be checked for their sync status and the
	// NetworkInfo gets updated
	NodeSyncInterval *IDuration `json:"nodeSyncInterval,omitempty" yaml:"nodeSyncInterval,omitempty" mapstructure:"nodeSyncInterval,omitempty"`

	// Nodes corresponds to the JSON schema field "nodes".
	Nodes []interface{} `json:"nodes,omitempty" yaml:"nodes,omitempty" mapstructure:"nodes,omitempty"`

	// Permanodes corresponds to the JSON schema field "permanodes".
	Permanodes []interface{} `json:"permanodes,omitempty" yaml:"permanodes,omitempty" mapstructure:"permanodes,omitempty"`

	// The amount of threads to be used for proof of work
	PowWorkerCount float64 `json:"powWorkerCount,omitempty" yaml:"powWorkerCount,omitempty" mapstructure:"powWorkerCount,omitempty"`

	// Node which will be tried first for all requests
	PrimaryNode interface{} `json:"primaryNode,omitempty" yaml:"primaryNode,omitempty" mapstructure:"primaryNode,omitempty"`

	// Node which will be tried first when using remote PoW, even before the
	// primary_node
	PrimaryPowNode interface{} `json:"primaryPowNode,omitempty" yaml:"primaryPowNode,omitempty" mapstructure:"primaryPowNode,omitempty"`

	// If node quorum is enabled. Will compare the responses from multiple nodes and
	// only returns the response if quorum_threshold of the nodes return the same one
	Quorum bool `json:"quorum,omitempty" yaml:"quorum,omitempty" mapstructure:"quorum,omitempty"`

	// % of nodes that have to return the same response so it gets accepted
	QuorumThreshold float64 `json:"quorumThreshold,omitempty" yaml:"quorumThreshold,omitempty" mapstructure:"quorumThreshold,omitempty"`

	// Timeout when sending a block that requires remote proof of work
	RemotePowTimeout *IDuration `json:"remotePowTimeout,omitempty" yaml:"remotePowTimeout,omitempty" mapstructure:"remotePowTimeout,omitempty"`
}

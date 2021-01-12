package config

// Kafka is a struct representing the Kafka configuration.
type Kafka struct {
	BootstrapServers *string `json:"bootstrapServers,omitempty"`
	GroupID          *string `json:"groupId,omitempty"`
	AutoOffsetReset  *string `json:"autoOffsetReset,omitempty"`
	SecurityProtocol *string `json:"securityProtocol,omitempty"`
	SASLMechanism    *string `json:"saslMechanism,omitempty"`
	SASLUsername     *string `json:"saslUsername,omitempty"`
	SASLPassword     *string `json:"saslPassword,omitempty"`
}

package model

import "errors"

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

// Validate validates the Kafka data.
func (k *Kafka) Validate() error {
	if k.BootstrapServers == nil || *k.BootstrapServers == "" {
		return errors.New("bootstrapServers is a required field")
	}

	if k.GroupID == nil || *k.GroupID == "" {
		return errors.New("groupId is a required field")
	}

	if k.AutoOffsetReset == nil || *k.AutoOffsetReset == "" {
		return errors.New("autoOffsetReset is a required field")
	}

	if k.SecurityProtocol == nil || *k.SecurityProtocol == "" {
		return errors.New("securityProtocol is a required field")
	}

	if k.SASLMechanism == nil || *k.SASLMechanism == "" {
		return errors.New("saslMechanism is a required field")
	}

	if k.SASLUsername == nil || *k.SASLUsername == "" {
		return errors.New("saslUsername is a required field")
	}

	if k.SASLPassword == nil || *k.SASLPassword == "" {
		return errors.New("saslPassword is a required field")
	}

	return nil
}

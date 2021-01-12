package listener

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/reecerussell/kafka-ui/config"
	"github.com/reecerussell/kafka-ui/logging"
	"github.com/reecerussell/kafka-ui/model"
)

const defaultPollTimeout = 250

type Listener interface {
	Start(p Processor) error
	Stop()
}

type listener struct {
	ctx    context.Context
	cancel func()
	topics []string
	cnf    *kafka.ConfigMap
}

func New(ctx context.Context, cnf *config.Config) Listener {
	ctx, cancel := context.WithCancel(ctx)
	topics := getTopics(cnf.Topics)
	kc := &kafka.ConfigMap{
		"bootstrap.servers": *cnf.Kafka.BootstrapServers,
		"group.id":          *cnf.Kafka.GroupID,
		"auto.offset.reset": *cnf.Kafka.AutoOffsetReset,
		"security.protocol": *cnf.Kafka.SecurityProtocol,
		"sasl.mechanism":    *cnf.Kafka.SASLMechanism,
		"sasl.username":     *cnf.Kafka.SASLUsername,
		"sasl.password":     *cnf.Kafka.SASLPassword,
	}

	return &listener{
		ctx:    ctx,
		cancel: cancel,
		topics: topics,
		cnf:    kc,
	}
}

func getTopics(topics []*model.Topic) []string {
	names := make([]string, len(topics))

	for i, t := range topics {
		names[i] = t.Name
	}

	return names
}

func (l *listener) Start(p Processor) error {
	c, err := kafka.NewConsumer(l.cnf)
	if err != nil {
		return fmt.Errorf("could not instantiate consumer: %v", err)
	}
	defer c.Close()

	c.SubscribeTopics(l.topics, nil)

	for {
		select {
		case <-l.ctx.Done():
			return nil
		default:
			ev := c.Poll(defaultPollTimeout)

			switch e := ev.(type) {
			case *kafka.Message:
				logging.Info("Message received from '%s'", *e.TopicPartition.Topic)
				p.Handle(e)
				break
			case *kafka.Error:
				logging.Error("Consumer error: %v", e.Error())
				break
			}

			break
		}
	}
}

func (l *listener) Stop() {
	l.cancel()
}

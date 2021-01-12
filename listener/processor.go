package listener

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/reecerussell/kafka-ui/ws"
)

type Processor interface {
	Handle(m *kafka.Message)
}

func NewProcessor(send func(m *ws.Message)) Processor {
	return &processor{send: send}
}

type processor struct {
	send func(m *ws.Message)
}

func (p *processor) Handle(m *kafka.Message) {
	topic := *m.TopicPartition.Topic
	data := map[string]interface{}{
		"topic": topic,
		"message": map[string]interface{}{
			"key":       string(m.Key),
			"value":     string(m.Value),
			"timestamp": m.Timestamp.UTC().Unix(),
		},
	}

	p.send(&ws.Message{
		Type:    "Message",
		Payload: data,
	})
}

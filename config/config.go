package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/reecerussell/kafka-ui/model"
)

type Config struct {
	filename string
	mu       sync.Mutex
	Topics   []*model.Topic `json:"topics"`
	Kafka    *Kafka         `json:"kafka,omitempty"`
}

func GetConfig(filename string) (*Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open config: %v", err)
	}

	var c Config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return nil, fmt.Errorf("could not read config: %v", err)
	}

	c.filename = filename
	c.mu = sync.Mutex{}
	return &c, nil
}

func (c *Config) Save() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.OpenFile(c.filename, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not open config: %v", err)
	}

	bytes, _ := json.Marshal(c)
	file.Truncate(0)
	file.Write(bytes)

	return nil
}

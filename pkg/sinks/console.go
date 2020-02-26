package sinks

import (
	"context"
	"encoding/json"
	"os"

	"github.com/opsgenie/kubernetes-event-exporter/pkg/kube"
)

type ConsoleConfig struct {
	Layout map[string]interface{} `yaml:"layout"`
}

func (c *ConsoleConfig) Validate() error {
	return nil
}

type Console struct {
	encoder *json.Encoder
	layout  map[string]interface{}
}

func NewConsole(cfg *ConsoleConfig) (Sink, error) {
	return &Console{
		encoder: json.NewEncoder(os.Stdout),
		layout:  cfg.Layout,
	}, nil
}

func (c *Console) Close() {
	// No-op
}

func (c *Console) Send(ctx context.Context, ev *kube.EnhancedEvent) error {
	if c.layout == nil {
		return c.encoder.Encode(ev)
	}

	res, err := convertLayoutTemplate(c.layout, ev)
	if err != nil {
		return err
	}

	return c.encoder.Encode(res)
}

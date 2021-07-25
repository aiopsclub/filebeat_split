package main

import (
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	"github.com/pkg/errors"
)

type Split struct {
	config Config
}

const (
	processorName = "split"
)

func newSplitProcessor(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()

	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	split := &Split{
		config: config,
	}
	return split, nil
}

func (s Split) Run(event *beat.Event) (*beat.Event, error) {
	return event, nil
}

func (s Split) String() string {
	return "split"
}

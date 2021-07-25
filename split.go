package main

import (
	"fmt"
	"strings"
	"time"

	"regexp"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/processors"
	"github.com/oliveagle/jsonpath"
	"github.com/pkg/errors"
)

type Split struct {
	config Config
}

const (
	processorName = "split"
)

var envTypeRegex, _ = regexp.Compile("/test/q/logs/collected/[-_a-zA-Z0-9]+/(prod|grey|prepare|beta|dev|simulation)/")

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

func (s Split) autoDetect(event *beat.Event, pathValue string) (*beat.Event, error) {
	fieldSplitResult := strings.Split(pathValue, s.config.Separator)
	switch s.config.KeyName {
	case "env_type":
		if envTypeRegex.Match([]byte(pathValue)) {
			event.PutValue(s.config.KeyName, fieldSplitResult[6])
		} else {
			event.PutValue(s.config.KeyName, "")
		}
	case "log_name":
		logName := fieldSplitResult[len(fieldSplitResult)-1]
		logNameSplit := strings.Split(logName, ".")
		event.PutValue(s.config.KeyName, strings.Join(logNameSplit[0:len(logNameSplit)-1], "."))
	case "app_code":
		event.PutValue(s.config.KeyName, fieldSplitResult[5])
	case "pod_name":
		if envTypeRegex.Match([]byte(pathValue)) {
			event.PutValue(s.config.KeyName, fieldSplitResult[8])
		} else {
			event.PutValue(s.config.KeyName, fieldSplitResult[6])
		}
	}
	return event, nil
}

func (s Split) Run(event *beat.Event) (*beat.Event, error) {
	if s.config.EnableTimeStamp {
		if _, err := event.GetValue("timestamp"); err == common.ErrKeyNotFound {
			timestamp := event.Timestamp.UnixNano() / 1e6
			event.PutValue("timestamp", timestamp)
		}
		return event, nil
	}
	if s.config.EnableSendTime {
		if _, err := event.GetValue("send_time"); err == common.ErrKeyNotFound {
			sendTime := time.Now().UnixNano() / 1e6
			event.PutValue("send_time", sendTime)
		}
		return event, nil
	}
	if s.config.FieldPath == "" {
		logp.Err("FieldPath is empty")
		return event, nil
	}

	res, err := jsonpath.JsonPathLookup(event.Fields, "$."+s.config.FieldPath)

	if err != nil {
		if !s.config.IgnoreError {
			return event, err
		}
		return event, nil
	}

	if s.config.Separator == "" {
		logp.Err("Separator is empty")
		return event, nil
	}
	if s.config.Mode != "manul" && s.config.Mode != "auto" {
		logp.Err("Mode is not correct")
		return event, nil
	}
	resStrs, ok := res.(string)
	if !ok {
		logp.Err(fmt.Sprintf("%v is not string", res))
		return event, nil
	}

	fieldSplitResult := strings.Split(resStrs, s.config.Separator)

	if s.config.Mode == "manul" {
		if len(fieldSplitResult) < s.config.Index+1 {
			logp.Err(fmt.Sprintf("index %d must lower than %d", s.config.Index, len(fieldSplitResult)))
			return event, nil
		}
		event.PutValue(s.config.KeyName, fieldSplitResult[s.config.Index])

	} else {
		return s.autoDetect(event, resStrs)
	}

	return event, nil
}

func (s Split) String() string {
	return "split"
}

package main

import (
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSplitRun(t *testing.T) {

	config := map[string]interface{}{
		"Index":           -1,
		"FieldPath":       "source_path",
		"Separator":       "/",
		"IgnoreError":     true,
		"KeyName":         "log_name",
		"EnvTypeEnable":   false,
		"EnableTimeStamp": true,
		"EnableSendTime":  true,
	}

	testConfig, err := common.NewConfigFrom(config)
	assert.NoError(t, err)
	p, err := newSplitProcessor(testConfig)
	require.NoError(t, err)

	event := &beat.Event{
		Fields:    common.MapStr{},
		Timestamp: time.Now(),
	}
	_, _ = event.Fields.Put("source_path",
		"/home/q/logs/collected/cm_cm_test_myorder/prod/prod-host-dep1-769d8c9cdb-7bgq8")

	newEvent, err := p.Run(event)
	assert.NoError(t, err)
	assert.NotNil(t, newEvent, "")

}

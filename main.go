package main

import (
	"github.com/elastic/beats/v7/libbeat/plugin"
	"github.com/elastic/beats/v7/libbeat/processors"
)

var Bundle = plugin.Bundle(
	processors.Plugin(processorName, newSplitProcessor),
)

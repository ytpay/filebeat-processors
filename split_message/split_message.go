package split_message

import (
	"fmt"
	"strings"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("split_message", New)
	jsprocessor.RegisterPlugin("SplitMessage", New)
}

type splitMessage struct {
	config
	log *logp.Logger
}

const (
	processorName = "split_message"
	logName       = "processor.split_message"
)

// New constructs a new split_message processor.
func New(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	p := &splitMessage{
		config: config,
		log:    logp.NewLogger(logName),
	}

	return p, nil
}

// For any errors returned, as long as the event is not empty, the event will still be sent
// refs libbeat/publisher/processing/processors.go:105
func (p *splitMessage) Run(event *beat.Event) (*beat.Event, error) {
	msg, err := event.GetValue(p.SourceField)
	if err != nil {
		if p.IgnoreFailure || (p.IgnoreMissing && errors.Cause(err) == common.ErrKeyNotFound) {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to get source field %s", p.SourceField)
	}

	message, ok := msg.(string)
	if !ok {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.New("failed to parse message")
	}

	fieldsValue := strings.Split(message, p.Delimiter)
	p.log.Debugf("message fields: %v", fieldsValue)
	if len(fieldsValue) < len(p.TargetFields) {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.Errorf("incorrect field length: %d, expected length: %d", len(fieldsValue), len(p.TargetFields))
	}

	for i, k := range p.TargetFields {
		_, err = event.PutValue(k, strings.TrimSpace(fieldsValue[i]))
		if err != nil {
			if p.IgnoreFailure {
				return event, nil
			}
			return event, errors.Wrapf(err, "failed to put event value key: %s, value: %s", k, strings.TrimSpace(fieldsValue[i]))
		}
	}

	if p.ProcessorsField != "" {
		_, err = event.PutValue(p.ProcessorsField, true)
		if err != nil {
			if p.IgnoreFailure {
				return event, nil
			}
			return event, errors.Wrapf(err, "failed to put event value key: %s, value: %t", p.ProcessorsField, true)
		}
	}

	return event, nil
}

func (p *splitMessage) String() string {
	return fmt.Sprintf("split_message=[source_field=%s,target_fields=%v,delimiter=%s]",
		p.SourceField, p.TargetFields, p.Delimiter)
}

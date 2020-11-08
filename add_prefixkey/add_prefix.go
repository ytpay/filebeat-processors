package add_prefix

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
	processors.RegisterPlugin("add_prefix", New)
	jsprocessor.RegisterPlugin("AddPrefix", New)
}

type addPrefix struct {
	config
	log *logp.Logger
}

const (
	processorName = "add_prefix"
	logName       = "processor.add_prefix"
)

// New constructs a new add_prefix processor.
func New(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	p := &addPrefix{
		config: config,
		log:    logp.NewLogger(logName),
	}

	return p, nil
}

func (p *addPrefix) Run(event *beat.Event) (*beat.Event, error) {
	sf, err := event.GetValue(p.SourceField)
	if err != nil {
		if p.IgnoreFailure || (p.IgnoreMissing && errors.Cause(err) == common.ErrKeyNotFound) {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to get source field %s", p.SourceField)
	}

	sfStr, ok := sf.(string)
	if !ok {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.New("failed to parse source field")
	}

	ss := strings.Split(sfStr, p.Delimiter)
	if len(ss) < 1 {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.New("failed to split source field")
	}

	p.log.Debugf("split source field: %s", ss)
	_, err = event.PutValue(p.TargetField, ss[0])
	if err != nil {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to put event value key: %s, value: %s", p.TargetField, ss[0])
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

func (p *addPrefix) String() string {
	return fmt.Sprintf("add_prefix=[source_field=%s,target_field=%s,delimiter=%s]", p.SourceField, p.TargetField, p.Delimiter)
}

package add_log_type

import (
	"fmt"
	"path/filepath"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("add_log_type", New)
	jsprocessor.RegisterPlugin("AddLogType", New)
}

type addLogType struct {
	config
	log *logp.Logger
}

const (
	processorName = "add_log_type"
	logName       = "processor.add_log_type"
)

// New constructs a new add_prefix processor.
func New(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	p := &addLogType{
		config: config,
		log:    logp.NewLogger(logName),
	}

	return p, nil
}

func (p *addLogType) Run(event *beat.Event) (*beat.Event, error) {
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

	fileExt := filepath.Ext(sfStr)
	if fileExt == "" {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.New("failed to get file ext")
	}
	p.log.Debugf("file ext: %s", fileExt)

	logType, ok := p.TypeMap[fileExt]
	if !ok {
		if fileExt == "" {
			if p.IgnoreFailure {
				return event, nil
			}
			return event, errors.New("failed to get file type from type map")
		}
	}
	_, err = event.PutValue(p.TargetField, logType)
	if err != nil {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to put event value key: %s, value: %s", p.TargetField, logType)
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

func (p *addLogType) String() string {
	return fmt.Sprintf("add_log_type=[source_field=%s,target_field=%s,type_map=%s]", p.SourceField, p.TargetField, p.TypeMap)
}

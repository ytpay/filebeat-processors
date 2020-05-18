package add_filename

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("add_filename", New)
	jsprocessor.RegisterPlugin("AddFileName", New)
}

type addFilename struct {
	config
	log *logp.Logger
}

const (
	processorName = "add_filename"
	logName       = "processor.add_filename"
)

// New constructs a new add_filename processor.
func New(cfg *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := cfg.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	p := &addFilename{
		config: config,
		log:    logp.NewLogger(logName),
	}

	return p, nil
}

func (p *addFilename) Run(event *beat.Event) (*beat.Event, error) {
	lfp, err := event.GetValue(p.SourceField)
	if err != nil {
		if p.IgnoreFailure || (p.IgnoreMissing && errors.Cause(err) == common.ErrKeyNotFound) {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to get source field %s", p.SourceField)
	}

	logFilePath, ok := lfp.(string)
	if !ok {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.New("failed to parse log file path")
	}

	logFileName := filepath.Base(logFilePath)
	if p.EnableTimestamp && p.TimestampFormat != "" {
		logFileName = strings.TrimSuffix(logFileName, filepath.Ext(logFileName)) + "-" + time.Now().Format(p.TimestampFormat) + filepath.Ext(logFileName)
	}
	p.log.Debugf("%s: %s", p.TargetField, logFileName)
	_, err = event.PutValue(p.TargetField, logFileName)
	if err != nil {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to put event value key: %s, value: %s", p.TargetField, logFileName)
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

func (p *addFilename) String() string {
	return fmt.Sprintf("add_filename=[source_field=%s,target_field=%s,timestamp_format=%s]",
		p.SourceField, p.TargetField, p.TimestampFormat)
}

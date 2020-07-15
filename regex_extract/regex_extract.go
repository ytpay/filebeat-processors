package regex_extract

import (
	"fmt"
	"regexp"
	"time"

	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	jsprocessor "github.com/elastic/beats/v7/libbeat/processors/script/javascript/module/processor"
)

func init() {
	processors.RegisterPlugin("regex_extract", New)
	jsprocessor.RegisterPlugin("RegexExtract", New)
}

type RegexExtract struct {
	config
	regex *regexp.Regexp
	log   *logp.Logger
}

const (
	processorName = "regex_extract"
	logName       = "processor.regex_extract"
)

func New(c *common.Config) (processors.Processor, error) {
	config := defaultConfig()
	if err := c.Unpack(&config); err != nil {
		return nil, errors.Wrapf(err, "fail to unpack the %v configuration", processorName)
	}

	regex, err := regexp.Compile(config.Regex)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to compile the regex %s", config.Regex)
	}

	return &RegexExtract{
		config: config,
		regex:  regex,
		log:    logp.NewLogger(logName),
	}, nil
}

func (p *RegexExtract) String() string {
	return fmt.Sprintf("%v=[regex=[%v]]", processorName, p.Regex)
}

func (p *RegexExtract) Run(event *beat.Event) (*beat.Event, error) {
	msg, err := event.GetValue(p.SourceField)
	if err != nil {
		if p.IgnoreFailure || (p.IgnoreMissing && errors.Cause(err) == common.ErrKeyNotFound) {
			return event, nil
		}
		return event, errors.Wrapf(err, "could not fetch value for key: %s", p.SourceField)
	}

	message, ok := msg.(string)
	if !ok {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.New("failed to parse message")
	}

	value := p.regex.FindString(message)
	if len(value) == 0 {
		if p.IgnoreMissing || p.IgnoreFailure {
			return event, nil
		} else {
			return event, errors.New("failed to parse message")
		}
	}

	// TODO: remove date format parsing?
	if p.TargetField == "timestamp" {
		timestamp, _ := time.Parse("2006-01-02 15:04:05.000", value)
		event.Timestamp = timestamp
	}

	_, err = event.PutValue(p.TargetField, value)
	if err != nil {
		if p.IgnoreFailure {
			return event, nil
		}
		return event, errors.Wrapf(err, "failed to put event value key: %s, value: %s", p.TargetField, value)
	}
	return event, nil
}

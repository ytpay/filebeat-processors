package regexextract

import (
	"fmt"
	"regexp"
	"time"

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
	config RegexExtractConfig
}

const processorName = "regexExtract"

func New(c *common.Config) (processors.Processor, error) {
	fc := defaultRegexConfig
	err := c.Unpack(&fc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack regex extract config")
	}

	return &RegexExtract{
		config: fc,
	}, nil
}

func (p *RegexExtract) String() string {
	return fmt.Sprintf("%v=[regex=[%v]]", processorName, p.config.Regex)
}

func (f *RegexExtract) Run(event *beat.Event) (*beat.Event, error) {
	r, _ := regexp.Compile(f.config.Regex)
	msg, err := event.GetValue(f.config.Field)

	if err != nil {
		return event, errors.Wrapf(err, "could not fetch value for key: %s", f.config.Field)
	}

	message, ok := msg.(string)

	if !ok {
		return event, errors.New("failed to parse message")
	}

	value := r.FindString(message)

	if len(value) == 0 {
		if f.config.IgnoreMissing {
			return event, nil
		} else {
			return event, errors.New("failed to parse message")
		}
	}

	if f.config.Target == "timestamp" {
		timestamp, _ := time.Parse("2006-01-02 15:04:05.000", value)
		event.Timestamp = timestamp
	}

	event.PutValue(f.config.Target, value)
	return event, nil
}

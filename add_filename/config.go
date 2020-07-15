package add_filename

import "regexp"

// config for add_filename processor.
type config struct {
	EnableTimestamp bool              `config:"enable_timestamp"`
	TimestampFormat string            `config:"timestamp_format"`
	ProcessorsField string            `config:"processors_field"`
	SourceField     string            `config:"source_field"`
	TargetField     string            `config:"target_field"`
	IgnoreMissing   bool              `config:"ignore_missing"`
	IgnoreFailure   bool              `config:"ignore_failure"`
	LogTypeField    string            `config:"log_type_field"`
	LogType         map[string]string `config:"log_type"`
	logTypeRex      map[string]*regexp.Regexp
}

func defaultConfig() config {
	return config{
		EnableTimestamp: false,
		TimestampFormat: "2006-01-02",
		ProcessorsField: "processors.add_filename",
		TargetField:     "filename",
		SourceField:     "log.file.path",
	}
}

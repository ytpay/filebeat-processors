package add_filename

// config for add_filename processor.
type config struct {
	EnableTimestamp bool   `config:"enable_timestamp"`
	TimestampFormat string `config:"timestamp_format"`
	ProcessorsField string `config:"processors_field"`
	SourceField     string `config:"source_field"`
	TargetField     string `config:"target_field"`
	IgnoreMissing   bool   `config:"ignore_missing"`
	IgnoreFailure   bool   `config:"ignore_failure"`
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

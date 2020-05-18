package split_message

// config for split_message processor.
type config struct {
	SourceField     string   `config:"source_field"`
	TargetFields    []string `config:"target_fields"`
	Delimiter       string   `config:"delimiter"`
	ProcessorsField string   `config:"processors_field"`
	IgnoreMissing   bool     `config:"ignore_missing"`
	IgnoreFailure   bool     `config:"ignore_failure"`
}

func defaultConfig() config {
	return config{
		SourceField:     "message",
		TargetFields:    []string{"split_message.timestamp", "split_message.hostname", "split_message.thread", "split_message.level", "split_message.logger", "split_message.sourceFile", "split_message.line", "split_message.serviceName", "split_message.traceId", "split_message.feTraceId", "split_message.message", "split_message.exception"},
		Delimiter:       "$$",
		ProcessorsField: "processors.split_message",
	}
}

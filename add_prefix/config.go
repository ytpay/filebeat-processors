package add_prefix

// config for add_prefix processor.
type config struct {
	Delimiter       string `config:"delimiter"`
	ProcessorsField string `config:"processors_field"`
	SourceField     string `config:"source_field"`
	TargetField     string `config:"target_field"`
	IgnoreMissing   bool   `config:"ignore_missing"`
	IgnoreFailure   bool   `config:"ignore_failure"`
}

func defaultConfig() config {
	return config{
		Delimiter:       ".",
		ProcessorsField: "processors.add_prefix",
		SourceField:     "filename",
		TargetField:     "log_prefix",
		IgnoreFailure:   true,
	}
}

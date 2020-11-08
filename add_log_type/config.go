package add_log_type

// config for add_log_type processor.
type config struct {
	ProcessorsField string            `config:"processors_field"`
	SourceField     string            `config:"source_field"`
	TargetField     string            `config:"target_field"`
	TypeMap         map[string]string `config:"type_map"`
	IgnoreMissing   bool              `config:"ignore_missing"`
	IgnoreFailure   bool              `config:"ignore_failure"`
}

func defaultConfig() config {
	return config{
		ProcessorsField: "processors.add_log_type",
		SourceField:     "filename",
		TargetField:     "add_log_type",
		TypeMap: map[string]string{
			".log":    "log",
			".sqllog": "sql",
			".error":  "error",
			".json":   "json",
		},
	}
}

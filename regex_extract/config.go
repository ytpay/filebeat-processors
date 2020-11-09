package regex_extract

type config struct {
	Regex         string `config:"regex"`
	SourceField   string `config:"source_field"`
	TargetField   string `config:"target_field"`
	IgnoreMissing bool   `config:"ignore_missing"`
	IgnoreFailure bool   `config:"ignore_failure"`
}

func defaultConfig() config {
	return config{
		Regex:         "[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}(?:.\\d{3}\\b)?",
		SourceField:   "message",
		TargetField:   "timestamp",
		IgnoreFailure: true,
	}
}

package regexextract

type RegexExtractConfig struct {
	Regex         string `config:"regex"`
	Field         string `config:"field"`
	Target        string `config:"target"`
	IgnoreMissing bool   `config:"ignoreMissing"`
}

var defaultRegexConfig = RegexExtractConfig{
	Regex:         "[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}(?:.\\d{3}\\b)?",
	Field:         "message",
	Target:        "timestamp",
	IgnoreMissing: true,
}
